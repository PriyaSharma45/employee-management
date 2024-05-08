package sql_client

import "employee-management/pkg/model"

// MockDBClient is a mock implementation of DBClient for testing purposes.
type MockDBClient struct {
	CreateFunc                     func(value interface{}) error
	ReadFunc                       func(id uint) (model.Employee, error)
	UpdateFunc                     func(id uint, updateField string, employeeVal interface{}) error
	DeleteFunc                     func(id uint) error
	GetEmployeesWithPaginationFunc func(page, limit int) ([]model.Employee, error)
}

// Create calls the mocked CreateFunc.
func (m *MockDBClient) Create(value interface{}) error {
	return m.CreateFunc(value)
}

// Read calls the mocked CreateFunc.
func (m *MockDBClient) Read(id uint) (model.Employee, error) {
	return m.ReadFunc(id)
}

func (m *MockDBClient) Update(id uint, updateField string, value interface{}) error {
	return m.UpdateFunc(id, updateField, value)
}

func (m *MockDBClient) Delete(id uint) error {
	return m.DeleteFunc(id)
}

func (m *MockDBClient) GetEmployeesWithPagination(page, limit int) ([]model.Employee, error) {
	return m.GetEmployeesWithPaginationFunc(page, limit)
}
