package todo

import (
	"fmt"
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
