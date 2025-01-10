package main

import "time"

// todo is a struct containing information about 
// a particular todo item.
type todo struct {
	Task      string
	Done      bool
	CreatedAt time.Time
	CompletedAt time.Time
}

// TodoList is a custom type with the underlying []todo data type.
// It contains various methods for creating and managing todos.
type TodoList []todo


// Add creates a new todo item and appends it to the list.
func (l *TodoList) Add(task string) {
	t := todo{
		Task: task,
		CreatedAt: time.Now(),
		Done: false,
	}

	*l = append(*l, t)
}