package main

import (
	"fmt"
	"go-ddd-layered/cmd/api"
	"go-ddd-layered/pkg/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := database.NewLocalPostgres()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	}).Methods("GET")

	api.RegisterUserEndpoints(r, db)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
