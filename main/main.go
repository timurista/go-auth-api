package main

import (
	// "database/sql"
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	// "log"
	// "net/http"
	// "strings"
	//
	// jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	// "github.com/lib/pq"
	// "golang.org/x/crypto/bcrypt"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	// fmt.Printf("hello, world\n")
	log.Println("Listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func signup(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("signup\n")
	w.Write([]byte("successfully called signup"))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully called login"))
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("protectedEndpoint\n")
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Printf("TokenVerifyMiddleWare\n")
	return nil
}
