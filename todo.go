package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// todo is a struct containing information about
// a particular todo item.
type todo struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// TodoList is a custom type with the underlying []todo data type.
// It contains various methods for creating and managing todos.
type TodoList []todo

// Add creates a new todo item and appends it to the list.
func (t *TodoList) Add(task string) {
	item := todo{
		Task:      task,
		CreatedAt: time.Now(),
		Done:      false,
	}

	*t = append(*t, item)
}

// Complete marks a todo item as completed.
//
// It sets Done field to true and CompletedAt field to the current time
func (t *TodoList) Complete(pos int) error {
	todoList := *t

	// Check if pos is within the range of todo list.
	if pos <= 0 || pos > len(todoList) {
		return fmt.Errorf("item %d does not exist", pos)
	}

	// Adjusting for 0-based index.
	todoList[pos-1].Done = true
	todoList[pos-1].CompletedAt = time.Now()

	return nil
}

// Delete a todo item from the todo list based on the
// provided position in the list.
func (t *TodoList) Delete(pos int) error {
	todoList := *t

	// Check if pos is within the range of todo list.
	if pos <= 0 || pos > len(todoList) {
		return fmt.Errorf("item %d does not exist", pos)
	}

	// Adjusting for 0-based index.
	*t = append(todoList[:pos-1], todoList[:pos]...)
	return nil
}

// Save encodes the TodoList as JSON and then saves it
// using the provided file name into the current working directory.
func (t *TodoList) Save(filename string) error {
	js, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}


// Get opens the provided file name, decodes the JSON data
// and parses it into the TodoList.
func (t *TodoList) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			return nil
		default:
			return err
		}
	}
	
	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, t)
}