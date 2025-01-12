package main

import (
	"fmt"
	"os"

	"github.com/hayohtee/todo"
)

const todoFileName = ".todo.json"

func main() {
	var todoList todo.TodoList

	if err := todoList.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}