package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ArgHandler contains a map of handlers
type ArgHandler struct {
	handlers map[string]HandlerEntry
	tl       *TaskList
}

// HandlerEntry contains a handler func and the command its registered to
type HandlerEntry struct {
	handler func(*TaskList, []string)
	command string
}

// Creates a new ArgHandler
func newArgHandler(tl *TaskList) *ArgHandler {
	return &ArgHandler{make(map[string]HandlerEntry), tl}
}

// Register a handler func to a command.
func (h *ArgHandler) registerHandler(command string, handlerFunc func(*TaskList, []string)) {
	if handlerFunc == nil {
		fmt.Println("Attempted to register nil handler")
		os.Exit(1)
	}
	if _, ok := h.handlers[command]; ok {
		fmt.Println(command, "handler already registered")
		os.Exit(1)
	}

	h.handlers[command] = HandlerEntry{handler: handlerFunc, command: command}
}

// Determine what handler to call for incoming arg slice and call it
func (h *ArgHandler) handle(args []string) {
	command := ""
	if len(args) > 1 {
		command = args[1]
	}

	var extraArgs []string
	if len(args) > 2 {
		extraArgs = args[2:]
	}

	if _, ok := h.handlers[command]; ok {
		h.handlers[command].handler(h.tl, extraArgs)
	} else {
		fmt.Println("No handler registered for:", command)
	}
}

// Handles an init command by creating a new task file if it does not exist.
func handleInit(tl *TaskList, args []string) {
	if err := createTaskFile(); err != nil {
		fmt.Println("A Task File already exists in the current directory.")
	}
	fmt.Println("Created a new Task File.")
}

// Handles a list command, prints out all or part of the current task list.
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

// Handles an add command, adds a task to the list.
func handleTaskAdd(tl *TaskList, args []string) {
	if err := tl.loadTasks(); err != nil {
		fmt.Println(err)
		return
	}
	if len(args) > 0 {
		tl.addTask(strings.Join(args[0:], " "))
		if err := tl.flushTasks(); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		displayHelp()
	}
}

// Handles a del command, deletes a task from the list.
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

// Handles an edit command, edits a task in the list.
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

// Handles a done command, completes a task in the list or uncompletes a already done task.
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

// Handles a store command, stores a done task or unstores a stored task.
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

// Handles a p command, does project specific operations.
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

// Handles a help command, displays help prompts.
func handleHelp(tl *TaskList, args []string) {
	displayHelp()
}
