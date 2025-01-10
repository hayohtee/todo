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

func TestComplete(t *testing.T) {
	var todoList todo.TodoList

	taskName := "New Task"
	todoList.Add(taskName)

	if todoList[0].Task != taskName {
		t.Errorf("expected: %q, got %q instead", taskName, todoList[0].Task)
	}

	if todoList[0].Done {
		t.Error("new task should not be completed")
	}

	todoList.Complete(1)

	if !todoList[0].Done {
		t.Error("new task should be completed")
	}
}

func TestDelete(t *testing.T) {
	var todoList todo.TodoList

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, v := range tasks {
		todoList.Add(v)
	}

	if todoList[0].Task != tasks[0] {
		t.Errorf("expected: %q, got %q instead", tasks[0], todoList[0].Task)
	}

	todoList.Delete(2)

	if len(todoList) != 2 {
		t.Errorf("expected list length: %d, got %d instead", 2, len(todoList))
	}

	if todoList[1].Task != tasks[2] {
		t.Errorf("expected: %q, got %q instead", tasks[2], todoList[1].Task)
	}
}
