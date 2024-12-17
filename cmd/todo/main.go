package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/hayohtee/todo"
)

const todoFileName = ".todo.json"

func main() {
	flag.Usage = func ()  {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed by Olamilekan Akintilebo\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright %d\n", time.Now().Year())
		fmt.Fprintln(flag.CommandLine.Output(), "Usage Information:")
		flag.PrintDefaults()
	}
	
	task := flag.String("task", "", "Task to be included in the todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	todoList := &todo.TodoList{}

	// Use the Get method to read todo items from file.
	if err := todoList.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {
	// List current todo items.
	case *list:
		for _, item := range *todoList {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}

	// Mark a todo item as completed
	case *complete > 0:
		// Complete the given item
		if err := todoList.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list.
		if err := todoList.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	// Add a new todo item
	case *task != "":
		todoList.Add(*task)

		// Save the new list
		if err := todoList.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	// Handle invalid flag
	default:
		fmt.Fprintln(os.Stderr, "invalid option")
		os.Exit(1)
	}
}
