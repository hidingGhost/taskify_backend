package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	DueDate string `json:"due_date"`
}

func GetTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT id, title, due_date FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.DueDate); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// CREATE Operation
func CreateTask(db *sql.DB, t Task) error {
	_, err := db.Exec("INSERT INTO tasks (title, due_date) VALUES ($1, $2)", t.Title, t.DueDate)
	return err
}
