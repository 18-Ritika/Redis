package main

import (
	handler "Training/Redis/Redis/internal/handlers/students"
	svc "Training/Redis/Redis/internal/services/students"
	"Training/Redis/Redis/internal/stores/students"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	store := students.New()
	service := svc.New(store)
	handler := handler.New(service)
	r := mux.NewRouter()
	r.HandleFunc("/students/{id}", handler.Get).Methods(http.MethodGet)
	r.HandleFunc("/students", handler.Post).Methods(http.MethodPost)
	r.HandleFunc("/students/{id}", handler.Delete).Methods(http.MethodDelete)
	http.ListenAndServe("localhost:8080", r)
}
