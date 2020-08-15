package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// create models
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// jwt models
type JWT struct {
	Token string `json:"token`
}

// error model
type Error struct {
	Message string `json:"message"`
}

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
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	fmt.Printf("password = %v\n", user.Password)
	fmt.Printf("password hash = %v\n", string(hash))

	token, _ := GenerateToken(user)
	fmt.Printf("token string = %v\n", token)

	//spew.Dump(token)
	//fmt.Println(tokenString)

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

func GenerateToken(user User) (string, error) {
	var err error
	var secret string = "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   "today",
	})

	var tokenString string
	tokenString, err = token.SignedString([]byte(secret))

	spew.Dump(token)

	return tokenString, err
}
