package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Status int

const (
	taskFile           = ".tasks.json"
	mainProject        = "main"
	Open        Status = iota
	Done
	Stored
)

type TaskList struct {
	Projects       map[string]Project `json:"projects,nilasempty"`
	CurrentProject string             `json:"currentproject"`
}

type Project struct {
	Tasks map[int]Task `json:"tasks,nilasempty"`
	Name  string       `json:"name"`
}

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

func newTaskList() *TaskList {
	tl := &TaskList{make(map[string]Project), mainProject}
	tl.Projects[mainProject] = Project{make(map[int]Task), mainProject}
	return tl
}

func (t *TaskList) currentList() Project {
	return t.Projects[t.CurrentProject]
}

func main() {
	// Handle init first since it needs to create the task file
	if len(os.Args) == 2 && os.Args[1] == "init" {
		if err := createTaskFile(); err != nil {
			fmt.Println("A Task File already exists in the current directory.")
			os.Exit(1)
		}
		fmt.Println("Created a new Task File.")
		os.Exit(0)
	}

	tl := newTaskList()
	if err := tl.loadTasks(); err != nil {
		// Handle PathError specifically as it indicates the file does not exist
		if _, ok := err.(*os.PathError); ok {
			fmt.Println("No task file found. Use 'tl init' to start a new task list here")
			os.Exit(0)
		}
		log.Fatal(err)
	}

	// Handle just "tl", which is an alias for "tl list open"
	if len(os.Args) < 2 {
		fmt.Println("\n" + tl.CurrentProject)
		tl.listTasks()
		return
	}

	switch os.Args[1] {
	case "list":
		if len(os.Args) > 2 {
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
		} else {
			fmt.Println("\n" + tl.CurrentProject)
			tl.listTasks()
			os.Exit(0)
		}
	case "add":
		if len(os.Args) > 2 {
			tl.addTask(strings.Join(os.Args[2:], " "))
		} else {
			displayHelp()
			os.Exit(0)
		}
	case "del":
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
	case "edit":
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
	case "done":
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
	case "store":
		if len(os.Args) > 2 {
			if os.Args[2] == "all" {
				tl.storeAll()
				break
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
	case "p":
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

	case "help":
		displayHelp()
		os.Exit(0)
	default:
		displayHelp()
		os.Exit(0)
	}

	if err := tl.flushTasks(); err != nil {
		log.Fatal(err)
	}
}
