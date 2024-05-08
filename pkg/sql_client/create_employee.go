package sql_client

import (
	"employee-management/pkg/model"
)

// Create inserts a new record into the database.
func (client *GORMClient) Create(value interface{}) error {
	db := client.DB.Create(value)
	return db.Error
}

// CreateNewEmployee adds a new employee to the database using the DBClient interface.
func CreateNewEmployee(client DBClient, employee model.Employee) error {
	return client.Create(&employee)
}
