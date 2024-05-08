package handler

import (
	"employee-management/pkg/model"
	"employee-management/pkg/sql_client"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func TestGetEmployeesWithPaginationHandler(t *testing.T) {
	mockDB := &sql_client.MockDBClient{
		GetEmployeesWithPaginationFunc: func(page, limit int) ([]model.Employee, error) {
			return []model.Employee{
				{EmployeeRecord: model.EmployeeRecord{Name: "John Doe", Position: "Software Engineer", Salary: 200000}},
				{EmployeeRecord: model.EmployeeRecord{Name: "Jane Doe", Position: "Project Manager", Salary: 100000}},
			}, nil
		},
	}

	handler := &RouteHandler{Engine: mockDB, mu: sync.Mutex{}}

	req, err := http.NewRequest("GET", "employees?page=1&limit=2", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler.GetEmployeesWithPaginationHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}

	expected := `"name":"John Doe"`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Expected response to contain %v, got %v", expected, rr.Body.String())
	}
}
