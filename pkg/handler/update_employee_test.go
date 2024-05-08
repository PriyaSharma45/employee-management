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

func TestUpdateEmployeeDHandler(t *testing.T) {
	// Mock database client
	mockDB := &sql_client.MockDBClient{
		UpdateFunc: func(id uint, updateField string, value interface{}) error {
			if id == 1 && updateField == "Position" {
				return nil
			}
			return fmt.Errorf("failed to update employee")
		},
	}

	handler := &RouteHandler{Engine: mockDB, mu: sync.Mutex{}}

	// Create test update data
	updateRequest := model.UpdateEmployee{
		ID:          1,
		UpdateField: "Position",
		Value:       "Lead Engineer",
	}
	updateJSON, _ := json.Marshal(updateRequest)

	req, err := http.NewRequest("POST", "/employee/update", bytes.NewBuffer(updateJSON))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.UpdateEmployeeDHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}

	expected := "successfully updated employee"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Expected response body to contain %v, got %v", expected, rr.Body.String())
	}

	// Test failure scenario
	updateRequest = model.UpdateEmployee{
		ID:          999,
		UpdateField: "Position",
		Value:       "Lead Engineer",
	}
	updateJSON, _ = json.Marshal(updateRequest)

	req, err = http.NewRequest("POST", "/employee/update", bytes.NewBuffer(updateJSON))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	handler.UpdateEmployeeDHandler(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Expected status code %v, got %v", http.StatusInternalServerError, status)
	}
}
