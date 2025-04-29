package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/LetsFocus/account-service/empdep/models"
)

func (h *handler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := h.app.GetEmployee(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsData)
}
func (h *handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		// Method not allowed
		http.Error(w, "Expecting POST method", http.StatusMethodNotAllowed)
		return
	}
	ctx := r.Context()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		// Bad request due to invalid or missing request body
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var employee models.Employee
	err = json.Unmarshal(data, &employee)
	if err != nil {
		// Bad request due to invalid JSON data
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	fmt.Println(employee)
	var createdEmployee models.Employee
	createdEmployee, err = h.app.CreateEmployee(ctx, employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsData, err := json.Marshal(createdEmployee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsData)
}

func (h *handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		// Method not allowed
		http.Error(w, "Expecting PUT method", http.StatusMethodNotAllowed)
		return
	}
	ctx := r.Context()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		// Bad request due to invalid or missing request body
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var employee models.Employee
	err = json.Unmarshal(data, &employee)
	if err != nil {
		// Bad request due to invalid JSON data
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	var updatedEmployee models.Employee
	updatedEmployee, err = h.app.UpdateEmployee(ctx, employee)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsData, err := json.Marshal(updatedEmployee)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsData)
}
func (h *handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		// Method not allowed
		http.Error(w, "Expecting DELETE method", http.StatusMethodNotAllowed)
		return
	}
	ctx := r.Context()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		// Bad request due to invalid or missing request body
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var employee models.Employee
	err = json.Unmarshal(data, &employee)
	if err != nil {
		// Bad request due to invalid JSON data
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	err = h.app.DeleteEmployee(ctx, employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
