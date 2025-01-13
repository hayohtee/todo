package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hayohtee/todo"
)

const todoFileName = ".todo.json"

func main() {
	// Parsing commandline flags
	task := flag.String("task", "", "Task to be included in the todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	var todoList todo.TodoList

	if err := todoList.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for _, item := range todoList {
			if !item.Done {
				fmt.Println(item)
			}
		}
	}
}
