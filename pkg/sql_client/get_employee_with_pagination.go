package sql_client

import "employee-management/pkg/model"

func (client *GORMClient) GetEmployeesWithPagination(page, limit int) ([]model.Employee, error) {
	var employees []model.Employee

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	result := client.DB.Offset(offset).Limit(limit).Find(&employees)

	return employees, result.Error
}
