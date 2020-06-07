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
	Done
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
	// Handle init first since it needs to create the task file
	if len(os.Args) == 2 && os.Args[1] == "init" {
		if err := createTaskFile(); err != nil {
			fmt.Println("A Task File already exists in the current directory.")
			os.Exit(1)
		}
		fmt.Println("Created a new Task File.")
		os.Exit(0)
	}

	tl := &TaskList{make(map[int]Task)}
	if err := tl.loadTasks(); err != nil {
		// Handle PathError specifically as it indicates the file does not exist
		if _, ok := err.(*os.PathError); ok {
			fmt.Println("No task file found. Use 'tl init' to start a new task list here")
			os.Exit(0)
		}
		log.Fatal(err)
	}

	// Handle just "tl", which is an alias for "tl list open"
	if len(os.Args) < 2 {
		tl.listTasks()
		return
	}

	switch os.Args[1] {
	case "list":
		tl.listTasks()
		os.Exit(0)
	case "add":
		if len(os.Args) > 2 {
			tl.addTask(os.Args[2])
		} else {
			displayHelp()
			os.Exit(0)
		}
	case "del":
		if len(os.Args) > 2 {
			id, err := strconv.Atoi(os.Args[2])
			logError(err)
			if err := tl.delTask(id); err != nil {
				fmt.Println(err)
			}
		} else {
			displayHelp()
			os.Exit(0)
		}
	case "edit":
		if len(os.Args) > 3 {
			id, err := strconv.Atoi(os.Args[2])
			logError(err)
			if err := tl.editTask(id, os.Args[3]); err != nil {
				fmt.Println(err)
			}
		} else {
			displayHelp()
			os.Exit(0)
		}
	case "done":
		if len(os.Args) > 2 {
			id, err := strconv.Atoi(os.Args[2])
			logError(err)
			if err := tl.doneTask(id); err != nil {
				fmt.Println(err)
			}
		} else {
			displayHelp()
			os.Exit(0)
		}
	case "reset":
		if len(os.Args) > 2 {
			id, err := strconv.Atoi(os.Args[2])
			logError(err)
			if err := tl.resetTask(id); err != nil {
				fmt.Println(err)
			}
		} else {
			displayHelp()
			os.Exit(0)
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
