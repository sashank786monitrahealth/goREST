package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var router *mux.Router = mux.NewRouter()
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	log.Println("Listen on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup invoked")
	w.Write([]byte("successfully called signup."))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login invoked")
	w.Write([]byte("successfully called login."))

}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protectedEndpoint invoked")

}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenVerifyMiddleware invoked")
	return nil
}
