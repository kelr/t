package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

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
		return tl, err
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
}

func (tl *TaskList) flushTasks() error {
	if _, err := os.Stat(taskList); err == nil {
		f, err := os.OpenFile(taskList, os.O_RDWR, 0644)
		logError(err)
		defer f.Close()

		err = json.NewEncoder(f).Encode(tl)
		if err == io.EOF {
			err = nil
		}
		return err
	} else if os.IsNotExist(err) {
		logError(err)
		return err
	}
	return nil
}
