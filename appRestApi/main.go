package main

import (
	"database/sql"
	//"./mysql-master"
	"net/http"
	"encoding/json"
	"fmt"
	"./mux-master"
	//"log"
	"os"
)

// Binding environmental variable via go package "os"
var getEnv = os.Getenv("LOGIN2")

// global variable for db connection
var db *sql.DB
var err error

type API struct {
	Message string "json: message"
}

type TestUser struct {
	ID int "json:id"
	Name string "json:name"
	Older int "json:older"
}

func helloApi (w http.ResponseWriter, r *http.Request) {
	helloApiMessage := "Hallo, hier ist deine Schnittstelle"
	message := API{helloApiMessage}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Irgendetwas ist schief gelaufen")
	}
	fmt.Fprintf(w, string(output))
}

func helloUser (w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	name := urlParams["user"]
	helloUserMessage := "Hallo " + name
	
	message := API{helloUserMessage}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Es lief etwas schief")
	}

	fmt.Fprintf(w, string(output))
}

func main() {

	// Open sql database with credentials from envVar
	// sql.Open(Driver_name, DB_user, DB_user_password from envVar, DB_adress, DB_name)
	/*
	db, err = sql.Open("mysql", "test_user:" + getEnv + "@tcp(localhost:3306)/testbase")
	if err != nil {
		log.Fatal(err)
	}
*/
	gorillaRoute := mux.NewRouter().StrictSlash(true)
	gorillaRoute.HandleFunc("/api", helloApi)
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", helloUser)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":3001", nil)

	defer db.Close()
}
