package main

import (
	"fmt"
)

// listProjects prints a list of available projects
func (t *TaskList) listProjects() {
	for name := range t.Projects {
		if name == t.CurrentProject {
			fmt.Println("*" + name)
		} else {
			fmt.Println(name)
		}
	}
}

// switchProject switches the current project.
func (t *TaskList) switchProject(name string) {
	if _, ok := t.Projects[name]; ok {
		t.CurrentProject = name
		fmt.Println("On project:", name)
	} else {
		fmt.Println("Project", name, "does not exist")
	}
}

// addProject adds a new project.
func (t *TaskList) addProject(name string) {
	if _, ok := t.Projects[name]; ok {
		fmt.Println("Project", name, "already exists")
	} else {
		t.Projects[name] = Project{make(map[int]Task), name}
		fmt.Println("Created new project", name)
	}
}

// editProject edits a current project's name
func (t *TaskList) editProject(name string, newName string) {
	if _, ok := t.Projects[name]; ok {
		if _, ok := t.Projects[newName]; ok {
			fmt.Println("Project", newName, "already exists")
			return
		}
		t.Projects[newName] = t.Projects[name]
		delete(t.Projects, name)
		if name == t.CurrentProject {
			t.CurrentProject = newName
		}
		fmt.Println("Project", name, "is now", newName)
	} else {
		fmt.Println("Project", name, "does not exist")
	}
}

// delProject deletes a project.
func (t *TaskList) delProject(name string) {
	if name == t.CurrentProject {
		fmt.Println("Cannot delete project while it is currently selected")
		return
	}

	if _, ok := t.Projects[name]; ok {
		delete(t.Projects, name)
		fmt.Println("Deleted project", name)
	} else {
		fmt.Println("Project", name, "does not exist")
	}
}
