package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/LetsFocus/account-service/empdep/models"
)

func (h *handler) GetDepatments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := h.app.GetDepatments(ctx)
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
func (h *handler) CreateDepartment(w http.ResponseWriter, r *http.Request) {

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

	var department models.Department
	err = json.Unmarshal(data, &department)
	if err != nil {
		// Bad request due to invalid JSON data
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	fmt.Print(department)
	var createdDepartment models.Department
	createdDepartment, err = h.app.CreateDepartment(ctx, department)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsData, err := json.Marshal(createdDepartment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsData)
}

func (h *handler) UpdateDepartment(w http.ResponseWriter, r *http.Request) {

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

	var department models.Department
	err = json.Unmarshal(data, &department)
	if err != nil {
		// Bad request due to invalid JSON data
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	var updatedDepartment models.Department
	updatedDepartment, err = h.app.UpdateDepartment(ctx, department)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsData, err := json.Marshal(updatedDepartment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsData)
}
func (h *handler) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
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

	var department models.Department
	err = json.Unmarshal(data, &department)
	if err != nil {
		// Bad request due to invalid JSON data
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	err = h.app.DeleteDepartment(ctx, department)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
