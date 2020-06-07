package main

import (
	"fmt"
	"log"
)

func displayHelp() {
	log.Println("help")
}

func (t *TaskList) listTasks() {
	for _, task := range t.Tasks {
		printTask(task)
	}
}

func printTask(task Task) {
	status := "[ ]"
	if task.Status != Open {
		status = "[X]"
	}
	fmt.Println(status, task.Id, "-", task.Description)
}

func (t *TaskList) addTask(task string) {
	newTask := &Task{
		Id:          len(t.Tasks),
		Description: task,
		Status:      Open,
	}
	t.Tasks = append(t.Tasks, *newTask)
}

func (t *TaskList) completeTask(Id int) {
	t.Tasks[Id].Status = Complete
}

func (t *TaskList) delTask() {
}
