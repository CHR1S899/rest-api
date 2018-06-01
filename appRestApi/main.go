package main

import (
	"database/sql"
	_"./mysql-master"
	"net/http"
	"encoding/json"
	"fmt"
	"./mux-master"
	"log"
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

type TodoArray struct {
	TodoList []Todo "json: array"
}

type Todo struct {
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

func createTodo (w http.ResponseWriter, r * http.Request) {
	NewTodo := Todo{}
	NewTodo.Task_Name = r.FormValue("task_name")
	output, err := json.Marshal(NewTodo)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Etwas ist schief gelaufen")
	}

	sql := "INSERT INTO todos_tbl (task_name) values ('" + NewTodo.Task_Name + "')"

	q, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(w, "Todo added: " + string(output))
	}
	fmt.Println(q)
}

func readTodos (w http.ResponseWriter, r *http.Request) {	
	
	/*
	var todoList TodoArray

	rows, err := db.Query("SELECT * FROM todos_tbl")
	
	defer rows.Close()
	
	// counting rows
	i := 0
	for rows.Next() {		
		i++
		var todo Todo
		err := rows.Scan(todo.Task_Id, todo.Task_Name, todo.Task_Status, todo.Task_Create, todo.Task_Update)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(todo)
		todoList.add(todo)
		
	}
	if i == 0 {
        log.Printf("Keine Todos vorhanden!")
        fmt.Fprintf(w, "Keine Todos vorhanden!")
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
	*/

	var (
		task_id string
		task_name string
		task_status string
		task_create string
		task_update string
	)	
	rows, err := db.Query("SELECT * FROM todos_tbl")
	
	defer rows.Close()
	
	// counting rows
	i := 0
	for rows.Next() {		
		i++
		err := rows.Scan(&task_id, &task_name, &task_status, &task_create, &task_update)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(task_id, task_name, task_status, task_create, task_update)
		
		resultString := "Task Id: " + task_id + " ,Task Name: " + task_name + " ,Task Status: " + task_status +
		" ,Task Create: " + task_create + " ,Task Update " + task_update + "\n"
		fmt.Fprintf(w, resultString)
	}
	if i == 0 {
        log.Printf("Keine Todos vorhanden!")
        fmt.Fprintf(w, "Keine Todos vorhanden!")
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
	}
	

}

func main() {

	// Open sql database with credentials from envVar
	// sql.Open(Driver_name, DB_user + DB_user_password from envVar, DB_adress, DB_name)	
	db, err = sql.Open("mysql", "todos_user:" + getEnv + "@tcp(localhost:3306)/todos_db")
	if err != nil {
		log.Fatal(err)
	}

	gorillaRoute := mux.NewRouter().StrictSlash(true)
	gorillaRoute.HandleFunc("/api", helloApi)
	gorillaRoute.HandleFunc("/api/todos/create", createTodo).Methods("GET")
	gorillaRoute.HandleFunc("/api/todos/read", readTodos).Methods("GET")
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":3001", nil)

	defer db.Close()
}
