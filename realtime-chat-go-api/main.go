package main

import (
	"fmt"
	"log"
	"net/http"
	"realtime-chat-go-api/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// documentation handler
	myRouter.HandleFunc("/", docPage)

	// users handler
	myRouter.HandleFunc("/users", handlers.UsersHandler).Methods("GET", "POST")
	myRouter.HandleFunc("/users/{id}", handlers.UserHandler).Methods("GET", "DELETE")

	// users message handler
	// myRouter.HandleFunc("/messages", handlers.MessagesHandler).Methods("GET", "POST")
	// myRouter.HandleFunc("/messages/{id}", handlers.MessagesHandler).Methods("GET", "DELETE")

	// JWT token handler
	myRouter.HandleFunc("/token/{username}", handlers.TokenHandler).Methods("GET")
	myRouter.HandleFunc("/token", handlers.RefreshTokenHandler).Methods("GET", "POST")

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8000", myRouter))

}

func docPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the My API docs page!")
}
