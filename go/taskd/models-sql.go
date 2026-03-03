package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			description TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func getTasks() ([]Task, error) {
	rows, err := db.Query("SELECT id, description FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Description); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func getTask(id int) (Task, error) {
	var t Task
	err := db.QueryRow("SELECT id, description FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Description)
	return t, err
}

func createTask(description string) (Task, error) {
	res, err := db.Exec("INSERT INTO tasks (description) VALUES (?)", description)
	if err != nil {
		return Task{}, err
	}

	id, _ := res.LastInsertId()
	return Task{ID: int(id), Description: description}, nil
}

func updateTask(id int, description string) error {
	_, err := db.Exec("UPDATE tasks SET description = ? WHERE id = ?", description, id)
	return err
}

func deleteTask(id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}