package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hayohtee/todo"
)

const todoFileName = ".todo.json"

func main() {
	todoList := &todo.TodoList{}

	// Use the Get method to read todo items from file.
	if err := todoList.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {
	// For no extra arguments, print the list
	case len(os.Args) == 1:
		for _, item := range *todoList {
			fmt.Println(item.Task)
		}

	// Concatenate all provided arguments with space and add to the todo list as item.
	default:
		// Concatenate all arguments with a space.
		item := strings.Join(os.Args[1:], " ")

		// Add the task.
		todoList.Add(item)

		// Save the new list.
		if err := todoList.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
