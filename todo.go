package todo

import (
	"fmt"
	"time"
)

// item is a struct that contains the fields for a todo item.
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// TodoList is a custom type with the underlying type of []item.
// It provide various method for creating and managing todos.
type TodoList []item

// Add is a method that add a new todo item to the todo list.
func (t *TodoList) Add(task string) {
	todoItem := item{
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
	}

	*t = append(*t, todoItem)
}

// Complete is a method that marks a todo item as completed
// by setting Done = true and CompletedTime to current time.
func (t *TodoList) Complete(i int) error {
	todoSlice := *t
	if i <= 0 || i > len(todoSlice) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjust the index for 0 based index
	todoSlice[i-1].Done = true
	todoSlice[i-1].CompletedAt = time.Now()

	return nil
}
