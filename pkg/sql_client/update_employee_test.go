package sql_client

import (
	"fmt"
	"testing"
)

func TestUpdateEmployeeDetails(t *testing.T) {
	mockDB := &MockDBClient{
		UpdateFunc: func(id uint, updateField string, value interface{}) error {
			if id == 1 && updateField == "Position" {
				return nil
			}
			return fmt.Errorf("failed to update employee")
		},
	}

	// Test successful update
	err := UpdateEmployeeDetails(mockDB, 1, "Position", "Lead Engineer")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test update failure
	err = UpdateEmployeeDetails(mockDB, 999, "Position", "Lead Engineer")
	if err == nil {
		t.Fatalf("Expected an error for non-existing employee, got nil")
	}
	expectedErr := "failed to update employee"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
	}
}
