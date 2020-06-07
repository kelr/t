package main

import (
	"log"
	"os"
	"strconv"
)

type status int

const (
	taskList        = "tasks"
	Open     status = iota
	Complete
	Archived
)

type Task struct {
	Id          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Status      status `json:"status,omitempty"`
}

type TaskList struct {
	Tasks []Task `json:"tasks,omitempty"`
}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func main() {
	tl = new(TaskList)
	if err := tl.loadTasks(); err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		tl.listTasks()
		return
	}

	switch os.Args[1] {
	case "list":
		tl.listTasks()
		os.Exit(0)
	case "add":
		tl.addTask(os.Args[2])
	case "del":
		tl.delTask()
	case "done":
		id, err := strconv.Atoi(os.Args[2])
		logError(err)
		tl.completeTask(id)
	case "help":
		displayHelp()
		os.Exit(0)
	default:
		displayHelp()
		os.Exit(0)
	}
	if err := tl.flushTasks(); err != nil {
		log.Fatal(err)
	}
}
