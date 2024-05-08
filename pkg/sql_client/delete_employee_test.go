package sql_client

import (
	"fmt"
	"testing"
)

func TestDeleteEmployeeDetails(t *testing.T) {
	mockDB := &MockDBClient{
		DeleteFunc: func(employeeID uint) error {
			if employeeID == 1 {
				return nil
			}
			return fmt.Errorf("employee not found")
		},
	}

	// Test successful delete
	err := DeleteEmployeeDetails(mockDB, 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test delete failure
	err = DeleteEmployeeDetails(mockDB, 999)
	if err == nil {
		t.Fatalf("Expected an error for non-existing employee, got nil")
	}
	expectedErr := "employee not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
	}
}
