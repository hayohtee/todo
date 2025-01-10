package todo_test

import (
	"testing"

	"github.com/hayohtee/todo"
)

func TestAdd(t *testing.T) {
	var todoList todo.TodoList

	taskName := "New Task"
	todoList.Add(taskName)

	if todoList[0].Task != taskName {
		t.Errorf("expected: %q, got %q instead", taskName, todoList[0].Task)
	}
}