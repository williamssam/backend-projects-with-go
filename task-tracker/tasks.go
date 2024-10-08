package main

import (
	"errors"
	"fmt"
	"log"
	"os"
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
	var nextID int

	for index := range *tasks {
		nextID = index + 1
	}

	task := Task{
		Id:          nextID,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	*tasks = append(*tasks, task)
	fmt.Printf("Task added successfully (ID: %d) \n", nextID)
}

func (tasks *Tasks) validateId(id int) error {
	if id < 0 && id >= len(*tasks) {
		err := errors.New("Task with id does not exist")
		fmt.Println(err)
		return err
	}

	return nil
}

func (tasks *Tasks) deleteTask(id int) error {
	t := *tasks

	err := t.validateId(id)
	if err != nil {
		return err
	}

	*tasks = append(t[:id], t[id+1:]...)
	fmt.Printf("Task (ID %d) deleted successfully \n", id)
	return nil
}

func (tasks *Tasks) markAsDone(id int) error {
	t := *tasks

	err := t.validateId(id)
	if err != nil {
		return err
	}

	status := t[id].Status
	if status == "done" {
		log.Fatalln("Todo already marked as done...")
	}

	t[id].Status = "done"
	fmt.Printf("Task (ID %d) marked as done \n", id)
	return nil
}

func (tasks *Tasks) markAsInProgress(id int) error {
	t := *tasks

	err := t.validateId(id)
	if err != nil {
		return err
	}

	status := t[id].Status
	if status == "in-progress" {
		log.Fatalln("Todo already marked as in-progress...")
	}

	if status == "done" {
		log.Fatalln("Todo already marked as done")
	}

	t[id].Status = "in-progress"
	fmt.Printf("Task (ID %d) marked as in progress \n", id)
	return nil
}

func (tasks *Tasks) listTasks() {
	t := *tasks
	var newTasks Tasks

	switch {
	case len(os.Args) <= 2:
		fmt.Println(t)
	case os.Args[2] == "done":
		for _, val := range t {
			if val.Status == "done" {
				newTasks = append(newTasks, val)
			}
		}
	case os.Args[2] == "in-progress":
		for _, val := range t {
			if val.Status == "in-progress" {
				newTasks = append(newTasks, val)
			}
		}
	case os.Args[2] == "todio":
		for _, val := range t {
			if val.Status == "todo" {
				newTasks = append(newTasks, val)
			}
		}
	default:
		fmt.Println(t)
	}

	fmt.Println(newTasks)
}
