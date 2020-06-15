package main

import (
	"os"
)

type Status int

const (
	Open Status = iota
	Done
	Stored
	mainProject = "main"
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
	handler.registerHandler("", handleList)
	handler.registerHandler("init", handleInit)
	handler.registerHandler("list", handleList)
	handler.registerHandler("add", handleTaskAdd)
	handler.registerHandler("del", handleTaskDel)
	handler.registerHandler("edit", handleTaskEdit)
	handler.registerHandler("done", handleTaskDone)
	handler.registerHandler("store", handleTaskStore)
	handler.registerHandler("p", handleProject)
	handler.registerHandler("help", handleHelp)
	handler.handle(os.Args)
}
