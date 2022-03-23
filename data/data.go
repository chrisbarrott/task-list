package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	// create connection
	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}

	// verify connection
	return db.Ping()
}

func CreateTable() {
	// create select query to get a task
	createTableSQL := `CREATE TABLE IF NOT EXISTS tasklist (
		"idTask" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"task" TEXT,
		"status" TEXT,
		"createdAt" TEXT
	  );`

	// update the query
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	// log results
	statement.Exec()
	log.Println("Task table created")
}

func InsertTask(task string, status string, createdAt string) {
	// create query
	insertNoteSQL := `INSERT INTO tasklist (task, status, createdAt)
	VALUES (?, ?, ?)`

	// update query
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}

	// execute pre compiles query
	_, err = statement.Exec(task, status, createdAt)
	if err != nil {
		log.Fatalln(err)
	}

	// log result
	log.Println("Inserted task into DB successfully")
}

func DisplayAllTasks() {
	// prop incase empty
	has_results := false

	// SQL select query
	row, err := db.Query("SELECT * FROM tasklist ORDER BY createdAt")
	if err != nil {
		log.Fatalln(err)
	}

	// row will close once when of function
	defer row.Close()

	// go through each row to get the value we need
	for row.Next() {
		// delclare found results
		has_results = true

		var idTask int
		var task string
		var status string
		var createdAt string

		row.Scan(&idTask, &task, &status, &createdAt)
		log.Println("[", status, "] - TASKID [", idTask, "]", task, "-", createdAt)
	}

	// check if task list is empty
	if !has_results {
		log.Fatalln("Nothing in tasklist! Create some tasks.")
	}
}

func UpdateStatusSQL(task string, status string) {
	var id *int
	var (
		idTask int
	)

	// query to sqlite
	rows, err := db.Query("SELECT idTask, task FROM tasklist WHERE task = ?", task)
	if err != nil {
		log.Fatal(err)
	}

	// go through each row to get the value we need
	for rows.Next() {
		err := rows.Scan(&idTask, &task)
		if err != nil {
			log.Fatal(err)
		}

		id = &idTask
	}

	// Check that the entry exists
	if id == nil {
		log.Fatal("Could not find task to update")
	}

	// query to update
	sqlUpdate := `UPDATE tasklist
		SET status = ? WHERE idTask = ?;`
	_, err = db.Exec(sqlUpdate, status, id)

	// error handling for execution
	if err != nil {
		log.Fatal(err)
	}

	// log result
	//log.Println("This should have updated task", task, "to", status)
}

func RemoveTaskSQL(task string) {
	// delcare variables
	statusCheck := "Completed"
	var id int

	// vartiables to get from database
	var (
		idTask int
		status string
	)

	// select to get row
	rows, err := db.Query("SELECT idTask, status FROM tasklist WHERE task = ?", task)
	if err != nil {
		log.Fatal(err)
	}

	// close connection
	defer rows.Close()

	// go through the rows to get the values we want
	for rows.Next() {
		err := rows.Scan(&idTask, &status)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Found task:", task)
	}

	// handle errors
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// set the id from the database to use later to delete
	id = idTask

	// check task exists
	if status == "" {
		log.Fatal("Cannot find task: ", task)
	}

	// check the status is completed before continuing
	if status != statusCheck {
		log.Fatal("Status of task ", task, " is not Completed, it is: ", status, ". Update the status to 'Completed' to proceed")
	}

	// query to update
	sqlUpdate := `DELETE FROM tasklist
		WHERE idTask = ?;`
	_, err = db.Exec(sqlUpdate, id)

	// error handling for execution
	if err != nil {
		log.Fatal(err)
	}

	// log result
	log.Println(task, "has been removed from task list")
}
