package main

import (
	"employee-management/pkg/handler"
	"employee-management/pkg/sql_client"
)

func main() {
	dbConnection := sql_client.GetSqlLiteClient()

	routerHandler := &handler.RouteHandler{
		Engine: dbConnection,
	}

	handler.GetRouterEngine(routerHandler)
}
