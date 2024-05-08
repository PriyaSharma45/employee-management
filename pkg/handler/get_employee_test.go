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

func TestGetEmployeeByIDHandler(t *testing.T) {
	// Mock database client
	mockDB := &sql_client.MockDBClient{
		ReadFunc: func(employeeID uint) (model.Employee, error) {
			if employeeID == 1 {
				return model.Employee{
					EmployeeRecord: model.EmployeeRecord{
						Name:     "John Doe",
						Position: "Software Engineer",
						Salary:   200000,
					},
				}, nil
			}
			return model.Employee{}, fmt.Errorf("employee not found")
		},
	}

	handler := &RouteHandler{Engine: mockDB, mu: sync.Mutex{}}

	// Create test employee data
	employeeID := model.EmployeeID{ID: 1}
	employeeIDJSON, _ := json.Marshal(employeeID)

	req, err := http.NewRequest("POST", "/employee/get", bytes.NewBuffer(employeeIDJSON))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.GetEmployeeByIDHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}

	expected := `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"John Doe","position":"Software Engineer","salary":200000}`
	if !strings.Contains(expected, rr.Body.String()) {
		t.Errorf("Expected response body %v, got %v", expected, rr.Body.String())
	}

	// Test failure scenario
	employeeID = model.EmployeeID{ID: 999}
	employeeIDJSON, _ = json.Marshal(employeeID)

	req, err = http.NewRequest("POST", "/employee/get", bytes.NewBuffer(employeeIDJSON))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	handler.GetEmployeeByIDHandler(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Expected status code %v, got %v", http.StatusInternalServerError, status)
	}
}
