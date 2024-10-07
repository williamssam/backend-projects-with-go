package main

import "fmt"

func main() {
	tasks := Tasks{}

	tasks.addTask("Samuel Williams")
	tasks.listTasks()
	fmt.Println("Welcome to task tracker")
}
