package main

import (
	"fmt"
	"log"
	"os"
)

type Status int

const (
	taskFile    = ".tasks.json"
	mainProject = "main"
)

const (
	Open Status = iota
	Done
	Stored
)

// TaskList contains a map of projects and the key to the current project
type TaskList struct {
	Projects       map[string]Project `json:"projects,nilasempty"`
	CurrentProject string             `json:"currentproject"`
}

// Project contains a map of Tasks and the project's name
type Project struct {
	Tasks map[int]Task `json:"tasks,nilasempty"`
	Name  string       `json:"name"`
}

// Task represents the current state of a task
type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

// Returns a new TaskList with the main project initialized
func newTaskList() *TaskList {
	tl := &TaskList{
		Projects:       make(map[string]Project),
		CurrentProject: mainProject,
	}
	tl.Projects[mainProject] = Project{
		Tasks: make(map[int]Task),
		Name:  mainProject,
	}
	return tl
}

// Get the current active project from the TaskList
func (t *TaskList) currentList() Project {
	return t.Projects[t.CurrentProject]
}

// Parse args and call the corresponding commands
func parseArgs(tl *TaskList) {
	// Handle just "tl", which is an alias for "tl list open"
	if len(os.Args) == 1 {
		tl.listTasks()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "list":
		handleList(tl)
	case "add":
		handleTaskAdd(tl)
	case "del":
		handleTaskDel(tl)
	case "edit":
		handleTaskEdit(tl)
	case "done":
		handleTaskDone(tl)
	case "store":
		handleTaskStore(tl)
	case "p":
		handleProject(tl)
	case "help":
		displayHelp()
		os.Exit(0)
	default:
		displayHelp()
		os.Exit(0)
	}
}

func main() {
	// Handle init first since it needs to create the task file
	if len(os.Args) == 2 && os.Args[1] == "init" {
		handleInit()
	}

	tl := newTaskList()
	if err := tl.loadTasks(); err != nil {
		// Handle PathError specifically as it indicates the file does not exist
		if _, ok := err.(*os.PathError); ok {
			fmt.Println("No task file found. Use 'tl init' to start a new task list here")
			os.Exit(0)
		}
		log.Fatal(err)
	}

	parseArgs(tl)

	// Flush the TaskList to the task file
	if err := tl.flushTasks(); err != nil {
		log.Fatal(err)
	}
}
