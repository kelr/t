package main

import (
	"fmt"
	"sort"
)

// Builds a list of tasks by status
func buildListStatus(currList Project, status Status) (out string) {
	out += "\n"
	var tasks []int
	for key := range currList.Tasks {
		if currList.Tasks[key].Status == status {
			tasks = append(tasks, key)
		}
		sort.Ints(tasks)
	}
	for _, id := range tasks {
		out += printTask(currList.Tasks[id])
	}
	return out
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
	currList := t.currentList()
	if _, ok := currList.Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	val := currList.Tasks[Id]
	if val.Status == Open {
		val.Status = Done
		currList.Tasks[Id] = val
		fmt.Println("Task", Id, "done")
	} else if val.Status == Done {
		val.Status = Open
		currList.Tasks[Id] = val
		fmt.Println("Task", Id, "open")
	} else {
		fmt.Println("Task", Id, "is archived.")
	}

	return nil
}

// Store a task.
func (t *TaskList) storeTask(Id int) error {
	currList := t.currentList()
	if _, ok := currList.Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	val := currList.Tasks[Id]
	if val.Status == Done {
		val.Status = Stored
		currList.Tasks[Id] = val
		fmt.Println("Task", Id, "stored")
	} else if val.Status == Open {
		val.Status = Open
		currList.Tasks[Id] = val
		fmt.Println("Cannot store open tasks!")
	} else {
		val.Status = Done
		currList.Tasks[Id] = val
		fmt.Println("Task", Id, "un-stored")
	}
	return nil
}

// Stores all done tasks
func (t *TaskList) storeAll() {
	count := 0
	currList := t.currentList()
	for Id := range currList.Tasks {
		val := currList.Tasks[Id]
		if val.Status == Done {
			val.Status = Stored
			currList.Tasks[Id] = val
			count++
		}
	}
	fmt.Println("Stored", count, "tasks")
}

// Deletes a task from the task list.
func (t *TaskList) delTask(Id int) error {
	currList := t.currentList()
	if _, ok := currList.Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	desc := currList.Tasks[Id].Description
	delete(currList.Tasks, Id)
	fmt.Println("Deleted Task", Id, "-", desc)
	return nil
}

// Edit a tasks description
func (t *TaskList) editTask(Id int, newDesc string) error {
	currList := t.currentList()
	if _, ok := currList.Tasks[Id]; !ok {
		return fmt.Errorf("Cannot find Task %d", Id)
	}
	val := currList.Tasks[Id]
	val.Description = newDesc
	currList.Tasks[Id] = val
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
func printTask(task Task) string {
	status := "[ ]"
	if task.Status != Open {
		status = "[X]"
	}
	return fmt.Sprintf("%s    %-5v %s\n", status, task.Id, task.Description)
}
