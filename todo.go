package todo

import (
	"encoding/json"
	"fmt"
	"os"
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

// Add a new todo item to the todo list.
func (t *TodoList) Add(task string) {
	todoItem := item{
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
	}

	*t = append(*t, todoItem)
}

// Complete marks a todo item as completed
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

// Delete  a todo item from the list.
func (t *TodoList) Delete(i int) error {
	todoSlice := *t

	if i <= 0 || i > len(todoSlice) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjusting the index for 0 based index.
	*t = append(todoSlice[:i-1], todoSlice[i:]...)

	return nil
}

// Save encode the TodoList as JSON and saves it to a
// disk using the provided filename.
func (t *TodoList) Save(filename string) error {
	js, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}