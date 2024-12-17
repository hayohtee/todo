package todo_test

import (
	"testing"

	"github.com/hayohtee/todo"
)

func TestAdd(t *testing.T) {
	todoList := todo.TodoList{}

	taskName := "New Task"
	todoList.Add(taskName)

	if todoList[0].Task != taskName {
		t.Errorf("expected %q, got %q instead", taskName, todoList[0].Task)
	}
}

func TestComplete(t *testing.T) {
	todoList := todo.TodoList{}

	taskName := "New Task"
	todoList.Add(taskName)

	if todoList[0].Task != taskName {
		t.Errorf("expected %q, got %q instead", taskName, todoList[0].Task)
	}

	todoList.Complete(1)

	if !todoList[0].Done {
		t.Errorf("new task should be completed")
	}
}
