package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

func main() {
	var router *mux.Router = mux.NewRouter()
	//fmt.Printf("type(router) = %T\n", router)
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("getBooks is called.")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("getBook is called.")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("addBook is called.")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("updateBook is called.")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("removeBook is called.")
}
