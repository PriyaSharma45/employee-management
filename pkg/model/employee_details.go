package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmployeeRecord
}

type EmployeeRecord struct {
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

type EmployeeID struct {
	ID uint `json:"id"`
}

type UpdateEmployee struct {
	ID          uint   `json:"id"`
	UpdateField string `json:"update_field"`
	Value       string `json:"value"`
}
