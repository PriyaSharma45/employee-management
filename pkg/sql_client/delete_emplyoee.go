package sql_client

import "employee-management/pkg/model"

// deletes the record from the database based on the ID.
func (client *GORMClient) Delete(employeeID uint) error {
	var employee model.Employee
	db := client.DB.Delete(&employee, employeeID)
	return db.Error
}

// DeleteEmployeeDetails deleted employee from the database using the DBClient interface.
func DeleteEmployeeDetails(client DBClient, employeeID uint) error {
	return client.Delete(employeeID)
}
