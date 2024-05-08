package sql_client

import (
	"employee-management/pkg/model"
	"fmt"
	"testing"
)

func TestGetEmployeeDetails(t *testing.T) {
	mockDB := &MockDBClient{
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

	// Test existing employee
	employee, err := GetEmployeeDetails(mockDB, 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if employee.Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", employee.Name)
	}
	if employee.Position != "Software Engineer" {
		t.Errorf("Expected position 'Software Engineer', got '%s'", employee.Position)
	}
	if employee.Salary != 200000 {
		t.Errorf("Expected department 'Engineering', got '%f'", employee.Salary)
	}

	// Test non-existing employee
	_, err = GetEmployeeDetails(mockDB, 999)
	if err == nil {
		t.Fatalf("Expected an error for non-existing employee, got nil")
	}
	expectedErr := "employee not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
	}
}
