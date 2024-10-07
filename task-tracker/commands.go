package main

import "flag"

type CommandFlags struct {
	Add            string
	Update         string
	Delete         bool
	MarkInProgress bool
	MarkDone       bool
	List           bool
	ListDone       bool
	ListTodos      bool
	ListInProgress bool
}

func NewCommand() *CommandFlags {
	cf := CommandFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add new task i.e. --add 'new task' ")
	flag.BoolVar(&cf.Delete, "delete", false, "Delete task using their id i.e. --delete 1")

	// Marking a task as in progress or done
	flag.BoolVar(&cf.MarkDone, "mark-done", false, "Mark task as done i.e. --mark-done 1")
	flag.BoolVar(&cf.MarkInProgress, "mark-in-progress", false, "Mark task as in progress i.e. --mark-in-progress 1")

	// Listing all tasks
	flag.BoolVar(&cf.List, "list", false, "List all tasks i.e. --list")

	// Listing all tasks by status

	return &cf

}
