package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

/*
	This function encodes the given map as a json string and
	writes it to the given ResponseWriter
*/
func writeResponse(w http.ResponseWriter, responseData map[string]interface{}) {

	responseString, err2 := json.Marshal(responseData)

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseString)
}

/*
	Global variables:
	store -- stores the cookie and session data
	db -- the interface to the database
*/

var store = sessions.NewCookieStore([]byte("this is not secure"))

var db *sql.DB

/*
	-Initialize the database interface
	-Route requests
*/
func main() {
	var err error
	db, err = sql.Open("mysql",
		"jack:jack@tcp(127.0.0.1:3306)/Calculator2")
	if err != nil {
		fmt.Println("Couldn't create database object")
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Couldn't ping db")
	}

	http.HandleFunc("/CalculatorAPI/StartSession", StartSession)
	http.HandleFunc("/CalculatorAPI/CreateAccount", CreateAccount)
	http.HandleFunc("/CalculatorAPI/CheckSession", CheckSession)
	http.HandleFunc("/CalculatorAPI/EndSession", EndSession)
	http.HandleFunc("/CalculatorAPI/AddCalculation", AddCalculation)
	http.HandleFunc("/CalculatorAPI/getLast10Calculations", getLast10Calculations)

	fs := http.FileServer(http.Dir("./static/dist"))
	http.Handle("/", fs)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
