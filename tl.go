package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
)

type status int

const (
	taskList        = "./tasks"
	Open     status = iota
	Complete
	Archived
)

var listCmd = flag.NewFlagSet("list", flag.ExitOnError)

type Task struct {
	Id          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Status      status `json:"status,omitempty"`
}

type TaskList struct {
	Tasks []Task `json:"tasks,omitempty"`
}

func displayHelp() {
	log.Println("help")
}

func (t *TaskList) listTasks() {
	log.Println("List cmd")
}

func (t *TaskList) addTask() {
	log.Println("Add cmd")
}

func (t *TaskList) delTask() {
	log.Println("Del cmd")
}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadTasks() (*TaskList, error) {
	tl := new(TaskList)
	if _, err := os.Stat(taskList); err == nil {
		f, err := os.Open(taskList)
		logError(err)
		defer f.Close()
		err = json.NewDecoder(f).Decode(tl)
		if err == io.EOF {
			err = nil
		}
	} else if os.IsNotExist(err) {
		f, err := os.Create(taskList)
		logError(err)
		defer f.Close()
		log.Println("New task list created")
		return tl, nil
	} else {
		// The file could exist but stat could fail due to perms
		return nil, err
	}
	return tl, nil
}

func main() {
	tl, err := loadTasks()
    if err != nil {
        logError(err)
    }
    if len(os.Args) < 2 {
		tl.listTasks()
		return
	}

	switch os.Args[1] {
	case "list":
		tl.listTasks()
	case "add":
		tl.addTask()
	case "del":
		tl.delTask()
	case "help":
		displayHelp()
		os.Exit(0)
	default:
		displayHelp()
		os.Exit(0)
	}

}
