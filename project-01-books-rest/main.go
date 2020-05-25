package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {
	var router *mux.Router = mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "Pride and Prejudice", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 2, Title: "Wuthering Heights", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 3, Title: "Middlemarch", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 4, Title: "Emma", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 5, Title: "Great Expectations", Author: "Charles Dickens", Year: "2010"},
	)

	//fmt.Printf("type(router) = %T\n", router)
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var params map[string]string = mux.Vars(r)

	log.Println(params)

	for _, book := range books {
		var bookParamID int64
		bookParamID, _ = strconv.ParseInt(params["id"], 10, 64)
		if book.ID == int(bookParamID) {
			json.NewEncoder(w).Encode(&book)
		}
	}

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
