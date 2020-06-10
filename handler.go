package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ArgHandler struct {
	handlers map[string]handlerEntry
	tl       *TaskList
}

type handlerEntry struct {
	handler func(*TaskList, []string)
	command string
}

func newArgHandler(tl *TaskList) *ArgHandler {
	return &ArgHandler{make(map[string]handlerEntry), tl}
}

func (h *ArgHandler) RegisterHandler(command string, handlerFunc func(*TaskList, []string)) {
	if handlerFunc == nil {
		fmt.Println("Attempted to register nil handler")
		os.Exit(1)
	}
	if _, ok := h.handlers[command]; ok {
		fmt.Println(command, "handler already registered")
		os.Exit(1)
	}

	h.handlers[command] = handlerEntry{handler: handlerFunc, command: command}
}

func (h *ArgHandler) Handle(args []string) {
	command := ""
	var extraArgs []string
	if len(args) > 1 {
		command = args[1]
	}
	if len(args) > 2 {
		extraArgs = args[2:]
	}

	if _, ok := h.handlers[command]; ok {
		h.handlers[command].handler(h.tl, extraArgs)
	} else {
		fmt.Println("No handler registered for:", command)
	}
}

func handleInit(tl *TaskList, args []string) {
	if err := createTaskFile(); err != nil {
		fmt.Println("A Task File already exists in the current directory.")
	}
	fmt.Println("Created a new Task File.")
}

func handleList(tl *TaskList, args []string) {
	if err := tl.loadTasks(); err != nil {
		fmt.Println(err)
		return
	}
	currList := tl.currentList()
	var output string
	if len(args) == 0 {
		fmt.Println("\n" + tl.CurrentProject)
		output += buildListStatus(currList, Open) + buildListStatus(currList, Done)
	} else {
		fmt.Println("\n" + tl.CurrentProject)
		if args[0] == "open" {
			output += buildListStatus(currList, Open)
		} else if args[0] == "done" {
			output += buildListStatus(currList, Done)
		} else if args[0] == "store" {
			output += buildListStatus(currList, Stored)
		} else if args[0] == "all" {
			output += buildListStatus(currList, Open) + buildListStatus(currList, Done) + buildListStatus(currList, Stored)
		} else {
			displayHelp()
			return
		}
	}
	fmt.Print(output)
}

func handleTaskAdd(tl *TaskList, args []string) {
	if err := tl.loadTasks(); err != nil {
		fmt.Println(err)
		return
	}
	if len(args) > 1 {
		tl.addTask(strings.Join(args[0:], " "))
		if err := tl.flushTasks(); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		displayHelp()
	}
}

func handleTaskDel(tl *TaskList, args []string) {
	if err := tl.loadTasks(); err != nil {
		fmt.Println(err)
		return
	}
	if len(args) == 1 {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid Task ID to delete")
			return
		}
		if err := tl.delTask(id); err != nil {
			fmt.Println(err)
		}
		if err := tl.flushTasks(); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		displayHelp()
	}
}

func handleTaskEdit(tl *TaskList, args []string) {
	if err := tl.loadTasks(); err != nil {
		fmt.Println(err)
		return
	}
	if len(args) > 1 {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid Task ID to edit")
			return
		}
		if err := tl.editTask(id, strings.Join(args[1:], " ")); err != nil {
			fmt.Println(err)
		}
		if err := tl.flushTasks(); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		displayHelp()
	}
}

func handleTaskDone(tl *TaskList, args []string) {
	if err := tl.loadTasks(); err != nil {
		fmt.Println(err)
		return
	}
	if len(args) == 1 {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid Task ID to set done")
			return
		}
		if err := tl.doneTask(id); err != nil {
			fmt.Println(err)
		}
		if err := tl.flushTasks(); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		displayHelp()
	}
}

func handleTaskStore(tl *TaskList, args []string) {
	if err := tl.loadTasks(); err != nil {
		fmt.Println(err)
		return
	}
	if len(args) == 1 {
		if args[0] == "all" {
			tl.storeAll()
		} else {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid Task ID to store")
				return
			}
			if err := tl.storeTask(id); err != nil {
				fmt.Println(err)
			}
		}
		if err := tl.flushTasks(); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		displayHelp()
	}
}

func handleProject(tl *TaskList, args []string) {
	if err := tl.loadTasks(); err != nil {
		fmt.Println(err)
		return
	}
	if len(args) == 0 {
		tl.listProjects()
	} else if len(args) == 1 {
		tl.switchProject(args[0])
		if err := tl.flushTasks(); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		switch args[0] {
		case "add":
			tl.addProject(args[1])
			if err := tl.flushTasks(); err != nil {
				fmt.Println(err)
				return
			}
		case "edit":
			if len(args) > 1 {
				tl.editProject(args[1], args[2])
				if err := tl.flushTasks(); err != nil {
					fmt.Println(err)
					return
				}
			} else {
				displayHelp()
			}
		case "del":
			tl.delProject(args[1])
			if err := tl.flushTasks(); err != nil {
				fmt.Println(err)
				return
			}
		default:
			displayHelp()
		}
	}
}

func handleHelp(tl *TaskList, args []string) {
	displayHelp()
}
