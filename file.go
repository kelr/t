package main

import (
	"encoding/json"
	"io"
	"os"
)

// Creates the task file. Returns a non nil error if the file already exists.
func createTaskFile() error {
	f, err := os.OpenFile(taskFile, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

// Opens the task file and read the contents into the TaskList.
func (tl *TaskList) loadTasks() error {
	f, err := os.OpenFile(taskFile, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = json.NewDecoder(f).Decode(tl); err == io.EOF {
		err = nil
	}
	return err
}

// Flush the contents of the TaskList to the task list file.
func (tl *TaskList) flushTasks() error {
	f, err := os.OpenFile(taskFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	if err = enc.Encode(tl); err == io.EOF {
		err = nil
	}
	return err
}
