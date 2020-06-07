package main

import (
	"fmt"
	"log"
	"sort"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

// Prints the current task list.
func (t *TaskList) listTasks() {
	if len(t.Tasks) == 0 {
		fmt.Println("Wowee no tasks")
	}

	var keys []int
	for key := range t.Tasks {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, id := range keys {
		printTask(t.Tasks[id])
	}
}

func (t *TaskList) getNextId() int {
	newId := 0
	for {
		if _, ok := t.Tasks[newId]; !ok {
			return newId
		}
		newId++
	}
}

// Adds a new task to the task list.
func (t *TaskList) addTask(task string) {
	newId := t.getNextId()
	newTask := &Task{
		Id:          newId,
		Description: task,
		Status:      Open,
	}
	t.Tasks[newId] = *newTask
	//t.Tasks = append(t.Tasks, *newTask)
	fmt.Println("Added Task", newTask.Id, "-", newTask.Description)
}

// Mark a task as complete.
func (t *TaskList) completeTask(Id int) error {
	if _, ok := t.Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	val := t.Tasks[Id]
	val.Status = Complete
	t.Tasks[Id] = val
	return nil
}

// Deletes a task from the task list.
func (t *TaskList) delTask(Id int) error {
	if _, ok := t.Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	delete(t.Tasks, Id)
	return nil
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
