package main

import (
	"fmt"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks []Task

func (tasks *Tasks) addTask(desc string) {
	task := Task{
		Id:          1,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	*tasks = append(*tasks, task)
}

func (tasks *Tasks) listTasks() {
	fmt.Println(tasks)
}
