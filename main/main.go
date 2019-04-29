package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	// "log"
	// "net/http"
	// "strings"
	//
	// jwt "github.com/dgrijalva/jwt-go"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int32  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

var db *sql.DB

func main() {

	// DB setup
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))

	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pgUrl)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pgUrl)

	err = db.Ping()

	router := mux.NewRouter()
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	// fmt.Printf("hello, world\n")
	log.Println("Listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func respondWithError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func responseJSON(w http.ResponseWriter, user User) {
	json.NewEncoder(w).Encode(user)
}

func GenerateToken(user User) (string, error) {
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})
	// spew.Dump(token)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, nil
}

func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is missing"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	// password hashing / salting
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	// fmt.Println(user)
	// fmt.Println("-------")
	if err != nil {
		log.Fatal(err)
		return
	}
	user.Password = string(hash)
	spew.Dump(user)

	stmt := "insert into users (email, password) values($1, $2) RETURNING id"
	err = db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		error.Message = "Server error"
		respondWithError(w, http.StatusInternalServerError, error)
		return
	}
	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	responseJSON(w, user)
}

func login(w http.ResponseWriter, r *http.Request) {

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	token, err := GenerateToken(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
	// w.Write([]byte("successfully called login"))
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("protectedEndpoint\n")
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Printf("TokenVerifyMiddleWare\n")
	return nil
}
