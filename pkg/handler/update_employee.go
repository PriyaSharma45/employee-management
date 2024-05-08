package handler

import (
	"employee-management/pkg/model"
	"employee-management/pkg/sql_client"
	"encoding/json"
	"log"
	"net/http"
)

func (client *RouteHandler) UpdateEmployeeDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update employee")

	client.mu.Lock()
	defer client.mu.Unlock()

	var request model.UpdateEmployee

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = sql_client.UpdateEmployeeDetails(client.Engine, request.ID, request.UpdateField, request.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("successfully updated employee"))
	w.WriteHeader(http.StatusOK)
}
