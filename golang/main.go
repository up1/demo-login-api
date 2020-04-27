package main

import (
	"database/sql"
	"demo/user"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var port = "12345"

func main() {
	// database, _ := sql.Open("sqlite3", "./users.db")
	// defer db.Close()
	// statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, username TEXT, password TEXT,firstname TEXT, lastname TEXT, email TEXT)")
	// defer statement.Close()
	// statement.Exec()
	// statement, _ = database.Prepare("INSERT INTO user (id, username, password, firstname, lastname, email) VALUES (?, ?, ?, ?, ?, ?)")
	// statement.Exec(1, "Somkiat", "PasswordSomkiat", "Somkiat", "Puisung", "somkiat@xxx.com")

	// var router *mux.Router
	router := mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.PathPrefix("/login").HandlerFunc(Login).Methods("POST")
	fmt.Println("Listening on port :12345")
	_ = http.ListenAndServe(":"+port, router)
}

// Login - Login to system with username and password
// URL : /api/v1/login
// Method: POST
// Body:
/*
 * {
 *	"username":"Somkiat",
 *	"password":"PasswordSomkiat"
 }
*/
// Output: JSON Encoded User object if created else JSON Encoded Exception.
func Login(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}
	defer db.Close()
	decoder := json.NewDecoder(r.Body)
	var loginRequest user.LoginRequest
	err = decoder.Decode(&loginRequest)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Some problem occurred with error %v", err))
		return
	}
	var id int
	var firstname string
	var lastname string
	var email string
	err = db.QueryRow("SELECT id, firstname, lastname, email FROM user WHERE username=? and password=?", loginRequest.UserName, loginRequest.Password).Scan(&id, &firstname, &lastname, &email)
	switch {
	case err == sql.ErrNoRows:
		respondWithError(w, http.StatusBadRequest, "No user found with the username="+loginRequest.UserName)
		return
	case err != nil:
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Some problem occurred with error %v", err))
		return
	default:
		var logedUser user.User
		logedUser.ID = id
		logedUser.FirstName = firstname
		logedUser.LastName = lastname
		logedUser.Email = email
		respondWithJSON(w, http.StatusOK, logedUser)
	}
}

// RespondWithError is called on an error to return info regarding error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Called for responses to encode and send json data
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//encode payload to json
	response, _ := json.Marshal(payload)
	// set headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}
