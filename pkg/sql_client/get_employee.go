package sql_client

import "employee-management/pkg/model"

// Reads the record from the database based on the ID.
func (client *GORMClient) Read(employeeID uint) (model.Employee, error) {
	var employee model.Employee
	db := client.DB.First(&employee, employeeID)
	if db.Error != nil {
		return employee, db.Error
	}
	return employee, nil
}

// GetEmployeeDetails gets an employee to the database using the DBClient interface.
func GetEmployeeDetails(client DBClient, employeeID uint) (model.Employee, error) {
	return client.Read(employeeID)
}
