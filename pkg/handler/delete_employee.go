package handler

import (
	"employee-management/pkg/model"
	"employee-management/pkg/sql_client"
	"encoding/json"
	"log"
	"net/http"
)

func (client *RouteHandler) DeleteEmployeeDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete an employee")

	client.mu.Lock()
	defer client.mu.Unlock()

	var request model.EmployeeID

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = sql_client.DeleteEmployeeDetails(client.Engine, request.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("successfully deleted employee"))
	w.WriteHeader(http.StatusOK)
}
