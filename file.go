package main

import (
	"encoding/json"
	"io"
	"os"
)

// Creates a new TaskList and populates it with the contents of the task list file
func (tl *TaskList) loadTasks() error {
	f, err := os.OpenFile(taskFile, os.O_RDWR|os.O_CREATE, 0755)
	logError(err)
	defer f.Close()

	if err = json.NewDecoder(f).Decode(tl); err == io.EOF {
		err = nil
	}
	return err
}

// Flush the contents of the TaskList to the task list file.
func (tl *TaskList) flushTasks() error {
	f, err := os.OpenFile(taskFile, os.O_RDWR|os.O_CREATE, 0755)
	logError(err)
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	if err = enc.Encode(tl); err == io.EOF {
		err = nil
	}
	return err
}
