package main

import "time"

// todo is a struct containing information about 
// a particular todo item.
type todo struct {
	Task      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
