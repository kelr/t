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

// Opens the task file and read the contents into the TaskList. File must exist.
func (tl *TaskList) loadTasks() error {
	f, err := os.OpenFile(taskFile, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	return tl.decodeFile(f)
}

// Decode the contents of a file into a TaskList struct
func (tl *TaskList) decodeFile(file *os.File) error {
	err := json.NewDecoder(file).Decode(tl)
	if err == io.EOF {
		err = nil
	}
	return err
}

// Flush the contents of the TaskList to the task list file.
func (tl *TaskList) flushTasks() error {
	f, err := os.OpenFile(taskFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	return tl.encodeFile(f)
}

// Encode the contents of a TaskList struct into a file
func (tl *TaskList) encodeFile(file *os.File) error {
	enc := json.NewEncoder(file)
	enc.SetIndent("", " ")
	err := enc.Encode(tl)
	if err == io.EOF {
		err = nil
	}
	return err
}
