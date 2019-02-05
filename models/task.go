package models

import (
	_ "database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanshirookazaki/echo-mvc/db"
)

type (
	Task struct {
		ID     int
		Userid int
		Task   string
		Status int
	}
	Tasks []Task

	T interface {
		GetTaskAll(int) Tasks
		GetTask(int) Tasks
		History(int) Tasks
		DeleteTask(int)
		FinishTask(int)
		EditTask(string, int)
		AddTask(int, string)
	}
)

func NewTask() *Task {
	return &Task{}
}

func (t *Task) GetTaskAll(userid int) Tasks {
	var tasks Tasks
	db, err := db.Conn()
	rows, err := db.Query("SELECT * FROM tasks WHERE Status = 0 and userid = " + strconv.Itoa(userid))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.Userid, &t.Task, &t.Status); err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, *t)
	}
	return tasks
}

func (t *Task) GetTask(id int) Tasks {
	var tasks Tasks
	db, err := db.Conn()
	rows, err := db.Query("SELECT * FROM tasks WHERE ID = " + strconv.Itoa(id))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.Userid, &t.Task, &t.Status); err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, *t)
	}
	return tasks
}

func (t *Task) History(userid int) Tasks {
	var tasks Tasks
	db, err := db.Conn()
	rows, err := db.Query("SELECT * FROM tasks WHERE Status = 1 and Userid =" + strconv.Itoa(userid))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.Userid, &t.Task, &t.Status); err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, *t)
	}
	return tasks
}

func (t *Task) DeleteTask(id int) {
	db, err := db.Conn()
	_, err = db.Query("DELETE FROM tasks WHERE ID =" + strconv.Itoa(id))
	if err != nil {
		panic(err.Error())
	}
}

func (t *Task) FinishTask(id int) {
	db, err := db.Conn()
	_, err = db.Query("UPDATE tasks SET Status = 1 WHERE ID =" + strconv.Itoa(id))
	if err != nil {
		panic(err.Error())
	}
}

func (t *Task) EditTask(task string, id int) {
	db, err := db.Conn()
	_, err = db.Query("UPDATE tasks SET Task = \"" + task + "\" WHERE ID = " + strconv.Itoa(id))
	if err != nil {
		panic(err.Error())
	}
}

func (t *Task) AddTask(userid int, task string) {
	db, err := db.Conn()
	_, err = db.Query("INSERT INTO tasks (Userid, Task, Status) VALUES ( " + strconv.Itoa(userid) + ", \"" + task + "\", 0)")
	if err != nil {
		panic(err.Error)
	}
}
