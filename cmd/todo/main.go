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
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	// Parsing commandline flags
	add := flag.Bool("add", false, "TAdd task to the todo list")
	list := flag.Bool("list", false, "List all tasks")
	del := flag.Int("del", 0, "Delete a todo item from the list")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	var todoList todo.TodoList

	if err := todoList.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		// List current to-do items.
		fmt.Println(&todoList)
	case *complete > 0:
		// Complete the given task.
		if err := todoList.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new todo list.
		if err := todoList.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		// Read the value of new task either from STDIN or command-line arguments.
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		todoList.Add(task)

		// Save the new list.
		if err := todoList.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del > 0:
		// Delete a todo based on the provided position.
		if err := todoList.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new todo list.
		if err := todoList.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "invalid option")
		os.Exit(1)
	}
}

// getTask decides where to get the description for a new task
// from (arguments or STDIN).
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
