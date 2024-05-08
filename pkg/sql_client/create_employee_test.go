package sql_client

import (
	"employee-management/pkg/model"
	"fmt"
	"testing"
)

func TestCreateNewEmployee(t *testing.T) {
	// Mock the Create function to simulate different scenarios
	mockClient := &MockDBClient{
		CreateFunc: func(value interface{}) error {
			return nil
		},
	}

	employee := model.Employee{
		EmployeeRecord: model.EmployeeRecord{Name: "Rachel",
			Position: "Software Engineer",
			Salary:   600000,
		},
	}

	err := CreateNewEmployee(mockClient, employee)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test failure scenario
	mockClient.CreateFunc = func(value interface{}) error {
		return fmt.Errorf("database error")
	}

	err = CreateNewEmployee(mockClient, employee)
	if err == nil {
		t.Fatalf("Expected an error, got none")
	}
}
