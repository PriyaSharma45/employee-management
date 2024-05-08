package sql_client

import "employee-management/pkg/model"

// Update the record from the database based on the ID.
func (client *GORMClient) Update(id uint, updateField string, value interface{}) error {
	var employee model.Employee
	db := client.DB.Model(&employee).Where("id", id).Update(updateField, value)
	return db.Error
}

// UpdateEmployeeDetails updates an employee to the database using the DBClient interface.
func UpdateEmployeeDetails(client DBClient, id uint, updateField string, value interface{}) error {
	return client.Update(id, updateField, value)
}
