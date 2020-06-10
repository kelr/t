package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleInit() {
	if err := createTaskFile(); err != nil {
		fmt.Println("A Task File already exists in the current directory.")
		os.Exit(1)
	}
	fmt.Println("Created a new Task File.")
	os.Exit(0)
}

func handleList(tl *TaskList) {
	if len(os.Args) == 2 {
		tl.listTasks()
		os.Exit(0)
	} else {
		fmt.Println("\n" + tl.CurrentProject)
		if os.Args[2] == "open" {
			tl.listStatus(Open)
		} else if os.Args[2] == "done" {
			tl.listStatus(Done)
		} else if os.Args[2] == "store" {
			tl.listStatus(Stored)
		} else if os.Args[2] == "all" {
			tl.listStatus(Open)
			tl.listStatus(Done)
			tl.listStatus(Stored)
		} else {
			displayHelp()
			os.Exit(0)
		}
	}
}

func handleTaskAdd(tl *TaskList) {
	if len(os.Args) > 2 {
		tl.addTask(strings.Join(os.Args[2:], " "))
	} else {
		displayHelp()
		os.Exit(0)
	}
}

func handleTaskDel(tl *TaskList) {
	if len(os.Args) > 2 {
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid Task ID to delete")
			os.Exit(1)
		}
		if err := tl.delTask(id); err != nil {
			fmt.Println(err)
		}
	} else {
		displayHelp()
		os.Exit(0)
	}
}

func handleTaskEdit(tl *TaskList) {
	if len(os.Args) > 3 {
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid Task ID to edit")
			os.Exit(1)
		}
		if err := tl.editTask(id, strings.Join(os.Args[3:], " ")); err != nil {
			fmt.Println(err)
		}
	} else {
		displayHelp()
		os.Exit(0)
	}
}

func handleTaskDone(tl *TaskList) {
	if len(os.Args) > 2 {
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid Task ID to set done")
			os.Exit(1)
		}
		if err := tl.doneTask(id); err != nil {
			fmt.Println(err)
		}
	} else {
		displayHelp()
		os.Exit(0)
	}
}

func handleTaskStore(tl *TaskList) {
	if len(os.Args) > 2 {
		if os.Args[2] == "all" {
			tl.storeAll()
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid Task ID to store")
			os.Exit(1)
		}
		if err := tl.storeTask(id); err != nil {
			fmt.Println(err)
		}
	} else {
		displayHelp()
		os.Exit(0)
	}
}

func handleProject(tl *TaskList) {
	if len(os.Args) == 2 {
		tl.listProjects()
	} else if len(os.Args) == 3 {
		tl.switchProject(os.Args[2])
	} else {
		switch os.Args[2] {
		case "add":
			tl.addProject(os.Args[3])
		case "edit":
			if len(os.Args) > 3 {
				tl.editProject(os.Args[3], os.Args[4])
			} else {
				displayHelp()
				os.Exit(0)
			}
		case "del":
			tl.delProject(os.Args[3])
		default:
			displayHelp()
			os.Exit(0)
		}
	}
}
