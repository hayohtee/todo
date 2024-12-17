package todo_test

import (
	"os"
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

func TestDelete(t *testing.T) {
	todoList := todo.TodoList{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, v := range tasks {
		todoList.Add(v)
	}

	if todoList[0].Task != tasks[0] {
		t.Errorf("expected %q, got %q instead", tasks[0], todoList[0].Task)
	}

	todoList.Delete(2)

	if len(todoList) != 2 {
		t.Errorf("expected list length %d, got %d instead", 2, len(todoList))
	}

	if todoList[1].Task != tasks[2] {
		t.Errorf("expected %q, got %q instead", tasks[2], todoList[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	todoList1 := todo.TodoList{}
	todoList2 := todo.TodoList{}

	taskName := "New Task"
	todoList1.Add(taskName)

	if todoList1[0].Task != taskName {
		t.Errorf("expected %q, got %q instead", taskName, todoList1[0].Task)
	}

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := todoList1.Save(tf.Name()); err != nil {
		t.Fatalf("error saving list to a file: %s", err)
	}

	if err := todoList2.Get(tf.Name()); err != nil {
		t.Fatalf("error getting list from file: %s", err)
	}

	if todoList1[0].Task != todoList2[0].Task {
		t.Errorf("task %q should match %q task", todoList1[0].Task, todoList2[0].Task)
	}
}
