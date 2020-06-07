package main

import (
	"fmt"
	"log"
)

// Prints the current task list.
func (t *TaskList) listTasks() {
	if len(t.Tasks) == 0 {
		fmt.Println("Wowee no tasks")
	}
	for _, task := range t.Tasks {
		printTask(task)
	}
}

// Adds a new task to the task list.
func (t *TaskList) addTask(task string) {
	newTask := &Task{
		Id:          len(t.Tasks),
		Description: task,
		Status:      Open,
	}
	t.Tasks = append(t.Tasks, *newTask)
}

// Mark a task as complete.
func (t *TaskList) completeTask(Id int) {
	t.Tasks[Id].Status = Complete
}

// Deletes a task from the task list.
func (t *TaskList) delTask() {
}

// Display the help menu.
func displayHelp() {
	log.Println("help")
}

// Prints out a single task.
func printTask(task Task) {
	status := "[ ]"
	if task.Status != Open {
		status = "[X]"
	}
	fmt.Println(status, task.Id, "-", task.Description)
}
