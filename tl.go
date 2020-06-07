package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type status int

const (
	taskFile        = "tasks"
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
	Tasks map[int]Task `json:"tasks,omitempty"`
}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func main() {
	tl := &TaskList{make(map[int]Task)}
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
		id, err := strconv.Atoi(os.Args[2])
		logError(err)
		if err := tl.delTask(id); err != nil {
			fmt.Println(err)
		}
	case "done":
		id, err := strconv.Atoi(os.Args[2])
		logError(err)
		if err := tl.completeTask(id); err != nil {
			fmt.Println(err)
		}
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
