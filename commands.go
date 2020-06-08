package main

import (
	"fmt"
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
	t.listStatus(Open)
	t.listStatus(Done)
}

// Prints tasks by status
func (t *TaskList) listStatus(status Status) {
	fmt.Println("")
	var tasks []int
	for key := range t.currentList().Tasks {
		if t.currentList().Tasks[key].Status == status {
			tasks = append(tasks, key)
		}
		sort.Ints(tasks)
	}
	for _, id := range tasks {
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

// Stores all done tasks
func (t *TaskList) storeAll() {
	count := 0
	for Id := range t.currentList().Tasks {
		val := t.currentList().Tasks[Id]
		if val.Status == Done {
			val.Status = Stored
			t.currentList().Tasks[Id] = val
			count++
		}
	}
	fmt.Println("Stored", count, "tasks")
}

// Deletes a task from the task list.
func (t *TaskList) delTask(Id int) error {
	if _, ok := t.currentList().Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	desc := t.currentList().Tasks[Id].Description
	delete(t.currentList().Tasks, Id)
	fmt.Println("Deleted Task", Id, "-", desc)
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
	fmt.Println("\nNew List:")
	fmt.Println("tl init\n")

	fmt.Println("Add Task:")
	fmt.Println("tl add New Task\n")

	fmt.Println("Show List:")
	fmt.Println("tl\n")

	fmt.Println("Complete Task:")
	fmt.Println("tl done 0\n")

	fmt.Println("Edit Task:")
	fmt.Println("tl edit 0 New Description\n")

	fmt.Println("Store Task:")
	fmt.Println("tl store 0\n")

	fmt.Println("Store All Done Tasks:")
	fmt.Println("tl store all\n")

	fmt.Println("Delete Task:")
	fmt.Println("tl del 0\n")

	fmt.Println("\nAdd Project:")
	fmt.Println("tl p add projectname\n")

	fmt.Println("List Projects:")
	fmt.Println("tl p\n")

	fmt.Println("Switch Current Project:")
	fmt.Println("tl p projectname\n")

	fmt.Println("Edit Project:")
	fmt.Println("tl p edit projectname newname\n")

	fmt.Println("Delete Project:")
	fmt.Println("tl p del projectname\n")
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
