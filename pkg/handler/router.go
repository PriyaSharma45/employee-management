package handler

import (
	"employee-management/pkg/sql_client"
	"net/http"
	"sync"
)

type RouteHandler struct {
	Engine sql_client.DBClient
	mu     sync.Mutex
}

func GetRouterEngine(rh *RouteHandler) {
	http.HandleFunc("/employee/create", rh.CreateNewEmployeeHandler)
	http.HandleFunc("/employee/get", rh.GetEmployeeByIDHandler)
	http.HandleFunc("/employee/update", rh.UpdateEmployeeDHandler)
	http.HandleFunc("/employee/delete", rh.DeleteEmployeeDHandler)
	http.HandleFunc("/employee", rh.GetEmployeesWithPaginationHandler)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
