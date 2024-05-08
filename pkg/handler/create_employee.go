package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"employee-management/pkg/model"
	"employee-management/pkg/sql_client"
)

func (client *RouteHandler) CreateNewEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Add a new employee")

	client.mu.Lock()
	defer client.mu.Unlock()

	var request model.EmployeeRecord

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = sql_client.CreateNewEmployee(client.Engine, model.Employee{EmployeeRecord: request})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("successfully entered employee"))
	w.WriteHeader(http.StatusOK)
}
