package todo

import "time"

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
