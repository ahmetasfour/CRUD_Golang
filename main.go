package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Starting server...")

	// Initialize the database
	initDatabase()
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	router := mux.NewRouter()

	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/users", createUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	//{____ Set up CORS options_________}
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},                   // Allow only your frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow these HTTP methods
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true, // Enable credentials, such as cookies or authentication headers

	})

	// Start the server on port 8000 with CORS handling
	fmt.Println("Server is running on port 8000...")
	if err := http.ListenAndServe(":8000", corsHandler.Handler(router)); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
