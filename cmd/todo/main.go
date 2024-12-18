package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/hayohtee/todo"
)

var todoFileName = ".todo.json"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed by Olamilekan Akintilebo\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright %d\n", time.Now().Year())
		fmt.Fprintln(flag.CommandLine.Output(), "Usage Information:")
		flag.PrintDefaults()
	}

	// Check if the user defined the ENV VAR for a custom file name.
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	add := flag.Bool("add", false, "Add task to the Todo list")
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
	// Add a new todo item
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		todoList.Add(t)

		if err := todoList.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	// List current todo items.
	case *list:
		fmt.Print(todoList)

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

	// Handle invalid flag
	default:
		fmt.Fprintln(os.Stderr, "invalid option")
		os.Exit(1)
	}
}

// getTask verify if any arguments were provided and then concatenate them into a
// string and return. If it doesn't, it default to reading from standard input (STDIN).
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", errors.New("task cannot be blank")
	}

	return s.Text(), nil
}
