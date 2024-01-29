package main

import (
	"net/http"

	"github.com/LetsFocus/account-service/empdep/handlers"
	"github.com/LetsFocus/account-service/empdep/services"
	"github.com/LetsFocus/account-service/empdep/store"
)

func main() {
	store := store.New()
	app := services.New(store)
	handler := handlers.New(app)
	http.HandleFunc("/getDep", handler.GetDepatments)
	http.HandleFunc("/createDep", handler.CreateDepartment)
	http.HandleFunc("/updateDep", handler.UpdateDepartment)
	http.HandleFunc("/deleteDep", handler.DeleteDepartment)
	http.HandleFunc("/createEmp", handler.CreateEmployee)
	http.HandleFunc("/getAll", handler.GetEmployee)
	http.HandleFunc("/updateEmp", handler.UpdateEmployee)
	http.HandleFunc("/deleteEmp", handler.DeleteEmployee)
	server := &http.Server{
		Addr: "127.0.0.1:8888",
	}
	server.ListenAndServe()

}
