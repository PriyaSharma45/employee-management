package handler

import (
	"employee-management/pkg/model"
	"employee-management/pkg/sql_client"
	"encoding/json"
	"log"
	"net/http"
)

func (client *RouteHandler) GetEmployeeByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Get employee value")

	var request model.EmployeeID

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	employeeDetails, err := sql_client.GetEmployeeDetails(client.Engine, request.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	detail, err := json.Marshal(employeeDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(detail)
	w.WriteHeader(http.StatusOK)
}
