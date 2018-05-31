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
	Task_Id int "json: task_id"
	Task_Name string "json: task_name"
	Task_Status string "json: task_status"
	Task_Create string "json: task_create"
	Task_Update string "json: task_Update"
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

func main() {

	//asdf
	// Open sql database with credentials from envVar
	// sql.Open(Driver_name, DB_user + DB_user_password from envVar, DB_adress, DB_name)
	/*
	db, err = sql.Open("mysql", "todos_user:" + getEnv + "@tcp(localhost:3306)/todos_db")
	if err != nil {
		log.Fatal(err)
	}
*/
	gorillaRoute := mux.NewRouter().StrictSlash(true)
	gorillaRoute.HandleFunc("/api", helloApi)
	gorillaRoute.HandleFunc("/api/todos}", helloUser)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":3001", nil)

	defer db.Close()
}
