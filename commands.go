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
	if len(t.currentList().Tasks) == 0 {
		fmt.Println("Task list is empty! Use 'tl add' to add a new task")
		return
	}
	fmt.Println("\n" + t.CurrentProject)

	var openTasks []int
	var doneTasks []int
	for key := range t.currentList().Tasks {
		if t.currentList().Tasks[key].Status == Open {
			openTasks = append(openTasks, key)
		} else if t.currentList().Tasks[key].Status == Done {
			doneTasks = append(doneTasks, key)
		}
	}
	sort.Ints(openTasks)
	sort.Ints(doneTasks)

	for _, id := range openTasks {
		printTask(t.currentList().Tasks[id])
	}
	fmt.Print("\n")
	for _, id := range doneTasks {
		printTask(t.currentList().Tasks[id])
	}
}

// Get the next available task Id.
func (t *TaskList) getNextId() int {
	newId := 0
	for {
		if _, ok := t.currentList().Tasks[newId]; !ok {
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
	t.currentList().Tasks[newId] = *newTask
	fmt.Println("Added Task", newTask.Id, "-", newTask.Description)
}

// Mark an open task as done.
func (t *TaskList) doneTask(Id int) error {
	if _, ok := t.currentList().Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	val := t.currentList().Tasks[Id]
	if val.Status == Open {
		val.Status = Done
		t.currentList().Tasks[Id] = val
		fmt.Println("Task", Id, "done")
	} else if val.Status == Done {
		val.Status = Open
		t.currentList().Tasks[Id] = val
		fmt.Println("Task", Id, "open")
	} else {
		fmt.Println("Task", Id, "is archived.")
	}

	return nil
}

// Store a task.
func (t *TaskList) storeTask(Id int) error {
	if _, ok := t.currentList().Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	val := t.currentList().Tasks[Id]
	if val.Status == Done {
		val.Status = Stored
		t.currentList().Tasks[Id] = val
		fmt.Println("Task", Id, "stored")
	} else if val.Status == Open {
		val.Status = Open
		t.currentList().Tasks[Id] = val
		fmt.Println("Cannot store open tasks!")
	} else {
		val.Status = Done
		t.currentList().Tasks[Id] = val
		fmt.Println("Task", Id, "un-stored")
	}

	return nil
}

// Deletes a task from the task list.
func (t *TaskList) delTask(Id int) error {
	if _, ok := t.currentList().Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	delete(t.currentList().Tasks, Id)
	fmt.Println("Deleted Task", Id)
	return nil
}

// Edit a tasks description
func (t *TaskList) editTask(Id int, newDesc string) error {
	if _, ok := t.currentList().Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	val := t.currentList().Tasks[Id]
	val.Description = newDesc
	t.currentList().Tasks[Id] = val
	fmt.Println("Edited Task", Id)
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
	fmt.Printf("%s    %-5v", status, task.Id)
	fmt.Println(task.Description)
}
