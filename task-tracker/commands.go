package main

import (
	"flag"
	"fmt"
)

type CommandFlags struct {
	Add            string
	Update         string
	Delete         int
	MarkInProgress int
	MarkDone       int
	List           bool
}

func createCommandFlags() *CommandFlags {
	cf := CommandFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add new task i.e. --add 'new task' ")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete task using their id i.e. --delete 1")

	// Marking a task as in progress or done
	flag.IntVar(&cf.MarkDone, "mark-done", -1, "Mark task as done i.e. --mark-done 1")
	flag.IntVar(&cf.MarkInProgress, "mark-in-progress", -1, "Mark task as in progress i.e. --mark-in-progress 1")

	// Listing all tasks
	flag.BoolVar(&cf.List, "list", false, "List all tasks i.e. --list")

	// Listing all tasks by status

	flag.Parse()

	return &cf
}

func (cmd *CommandFlags) execute(tasks *Tasks) {
	switch {
	case cmd.Add != "":
		tasks.addTask(cmd.Add)
	case cmd.MarkDone != -1:
		tasks.markAsDone(cmd.MarkDone)
	case cmd.MarkInProgress != -1:
		tasks.markAsInProgress(cmd.MarkInProgress)
	case cmd.List:
		tasks.listTasks()
	case cmd.Delete != -1:
		tasks.deleteTask(cmd.Delete)
	default:
		fmt.Println("Invalid command, plese use -h to get all commands")
	}
}
