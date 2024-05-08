package handler

import (
	"bytes"
	"employee-management/pkg/model"
	"employee-management/pkg/sql_client"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateNewEmployeeHandler(t *testing.T) {
	// Mock database client
	mockDB := &sql_client.MockDBClient{
		CreateFunc: func(value interface{}) error {
			// Simulate successful creation
			return nil
		},
	}

	handler := &RouteHandler{Engine: mockDB}

	// Create test employee data
	employeeRecord := model.EmployeeRecord{
		Name:     "Rachel",
		Position: "Engineer",
		Salary:   100000,
	}
	employeeJSON, _ := json.Marshal(employeeRecord)

	req, err := http.NewRequest("POST", "employees/create", bytes.NewBuffer(employeeJSON))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.CreateNewEmployeeHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}

	expected := "successfully entered employee"
	if rr.Body.String() != expected {
		t.Errorf("Expected response body %v, got %v", expected, rr.Body.String())
	}

	// Test failure scenario
	mockDB.CreateFunc = func(value interface{}) error {
		return fmt.Errorf("database error")
	}

	req, err = http.NewRequest("POST", "employees/create", bytes.NewBuffer(employeeJSON))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	rr = httptest.NewRecorder()
	handler.CreateNewEmployeeHandler(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Expected status code %v, got %v", http.StatusInternalServerError, status)
	}
}
