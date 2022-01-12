package main

import (
	"fmt"
	// "log"
	"net/http"
	// "log"
	// "net/http"
	// "github.com/alikulovDev/go-getting-started/db"
	// "github.com/gorilla/mux"
	// "gitlab.com/azamat.aliqulov/go_rest_api/-/blob/main/db.go"
	// "gitlab.com/azamat.aliqulov/go_rest_api/-/blob/main/quizController.go"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello world")
}

// func handleRequest() {
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("/", helloWorld).Methods("GET")
// 	// myRouter.HandleFunc("/users", AllUsers).Methods("GEt")
// 	fmt.Println("Hello GO exstensions")
// 	fmt.Println("Server at 8082")
// 	log.Fatal(http.ListenAndServe(":8082", myRouter))
// }

func main() {
	fmt.Println("Go ORM Tutorial")
	// db_cennetion.sayHello()
	// db_connection()
	// handleRequest()
}
