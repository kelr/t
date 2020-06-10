package main

import (
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

func main() {
	handler := newArgHandler(newTaskList())
	handler.RegisterHandler("", handleList)
	handler.RegisterHandler("init", handleInit)
	handler.RegisterHandler("list", handleList)
	handler.RegisterHandler("add", handleTaskAdd)
	handler.RegisterHandler("del", handleTaskDel)
	handler.RegisterHandler("edit", handleTaskEdit)
	handler.RegisterHandler("done", handleTaskDone)
	handler.RegisterHandler("store", handleTaskStore)
	handler.RegisterHandler("p", handleProject)
	handler.RegisterHandler("help", handleHelp)
	handler.Handle(os.Args)
}
