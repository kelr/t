package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type status int

const (
	taskFile           = ".tasks.json"
	mainProject        = "main"
	Open        status = iota
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
	Status      status `json:"status"`
}

func newTaskList() *TaskList {
	tl := &TaskList{make(map[string]Project), mainProject}
	tl.Projects[mainProject] = Project{make(map[int]Task), mainProject}
	return tl
}

func (t *TaskList) currentList() Project {
	return t.Projects[t.CurrentProject]
}

func (t *TaskList) switchProject(name string) {
	if _, ok := t.Projects[name]; ok {
		t.CurrentProject = name
		fmt.Println("On project:", name)
	} else {
		fmt.Println("Project", name, "does not exist")
	}
}

func (t *TaskList) addProject(name string) {
	if _, ok := t.Projects[name]; ok {
		fmt.Println("Project", name, "already exists")
	} else {
		t.Projects[name] = Project{make(map[int]Task), name}
		fmt.Println("Created new project", name)
	}
}

func (t *TaskList) editProject(name string, newName string) {
	if _, ok := t.Projects[name]; ok {
		t.Projects[newName] = t.Projects[name]
		delete(t.Projects, name)
		fmt.Println("Project", name, "is now", newName)
	} else {
		fmt.Println("Project", name, "does not exist")
	}
}

func (t *TaskList) delProject(name string) {
	if _, ok := t.Projects[name]; ok {
		delete(t.Projects, name)
		fmt.Println("Deleted project", name)
	} else {
		fmt.Println("Project", name, "does not exist")
	}
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
		tl.listTasks()
		return
	}

	switch os.Args[1] {
	case "list":
		tl.listTasks()
		os.Exit(0)
	case "add":
		if len(os.Args) > 2 {
			tl.addTask(os.Args[2])
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
			if err := tl.editTask(id, os.Args[3]); err != nil {
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
	case "switch":
		if len(os.Args) > 2 {
			tl.switchProject(os.Args[2])
		} else {
			displayHelp()
			os.Exit(0)
		}

	case "addproj":
		if len(os.Args) > 2 {
			tl.addProject(os.Args[2])
		} else {
			displayHelp()
			os.Exit(0)
		}

	case "editproj":
		if len(os.Args) > 3 {
			tl.editProject(os.Args[2], os.Args[3])
		} else {
			displayHelp()
			os.Exit(0)
		}

	case "delproj":
		if len(os.Args) > 2 {
			tl.delProject(os.Args[2])
		} else {
			displayHelp()
			os.Exit(0)
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
