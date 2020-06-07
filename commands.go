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
		fmt.Println("Task list is empty! Use 'tl add' to add a new task")
	}
	fmt.Println("Open =============================")
	var openTasks []int
	var doneTasks []int
	for key := range t.Tasks {
		if t.Tasks[key].Status == Open {
			openTasks = append(openTasks, key)
		} else if t.Tasks[key].Status == Done {
			doneTasks = append(doneTasks, key)
		}
	}
	sort.Ints(openTasks)
	sort.Ints(doneTasks)

	for _, id := range openTasks {
		printTask(t.Tasks[id])
	}
	if len(doneTasks) > 0 {
		fmt.Println("Done =============================")
	}
	for _, id := range doneTasks {
		printTask(t.Tasks[id])
	}
}

// Get the next available task Id.
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

// Mark an open task as done.
func (t *TaskList) doneTask(Id int) error {
	if _, ok := t.Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	val := t.Tasks[Id]
	val.Status = Done
	t.Tasks[Id] = val
	return nil
}

// Mark a done task as open again.
func (t *TaskList) resetTask(Id int) error {
	if _, ok := t.Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	val := t.Tasks[Id]
	val.Status = Open
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

// Edit a tasks description
func (t *TaskList) editTask(Id int, newDesc string) error {
	if _, ok := t.Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	val := t.Tasks[Id]
	val.Description = newDesc
	t.Tasks[Id] = val
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
