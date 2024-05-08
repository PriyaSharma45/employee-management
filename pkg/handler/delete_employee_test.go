package handler

import (
	"bytes"
	"employee-management/pkg/model"
	"employee-management/pkg/sql_client"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func TestDeleteEmployeeDHandler(t *testing.T) {
	// Mock database client
	mockDB := &sql_client.MockDBClient{
		DeleteFunc: func(employeeID uint) error {
			if employeeID == 1 {
				return nil
			}
			return fmt.Errorf("employee not found")
		},
	}

	handler := &RouteHandler{Engine: mockDB, mu: sync.Mutex{}}

	// Create test delete data
	deleteRequest := model.EmployeeID{
		ID: 1,
	}
	deleteJSON, _ := json.Marshal(deleteRequest)

	req, err := http.NewRequest("POST", "employees/delete", bytes.NewBuffer(deleteJSON))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.DeleteEmployeeDHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}

	expected := "successfully deleted employee"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Expected response body to contain %v, got %v", expected, rr.Body.String())
	}

	// Test failure scenario
	deleteRequest = model.EmployeeID{
		ID: 999,
	}
	deleteJSON, _ = json.Marshal(deleteRequest)

	req, err = http.NewRequest("POST", "employees/delete", bytes.NewBuffer(deleteJSON))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	handler.DeleteEmployeeDHandler(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Expected status code %v, got %v", http.StatusInternalServerError, status)
	}
}
