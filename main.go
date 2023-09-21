package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unawaretub86/task-rest/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.Ping)
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/create-task", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/task/{id}", handlers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/update-task/{id}", handlers.UpdateTask).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":3001", router))
}
