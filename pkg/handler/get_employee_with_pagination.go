package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (h *RouteHandler) GetEmployeesWithPaginationHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Get employees with pagination")

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	employees, err := h.Engine.GetEmployeesWithPagination(page, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	detail, err := json.Marshal(employees)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(detail)
	w.WriteHeader(http.StatusOK)
}
