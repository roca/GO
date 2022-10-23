package todo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %s, got %s instead", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %s, got %s instead", taskName, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("New Task should not be completed")
	}
	l.Complete(1)
	if !l[0].Done {
		t.Errorf("New Task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{"Task 1", "Task 2", "Task 3"}
	for _, task := range tasks {
		l.Add(task)
	}
	if l[0].Task != tasks[0] {
		t.Errorf("Expected %s, got %s instead", tasks[0], l[0].Task)
	}
	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("Expected 2 items, got %d instead", len(l))
	}
	if l[1].Task != tasks[2] {
		t.Errorf("Expected %s, got %s instead", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)
	if l1[0].Task != taskName {
		t.Errorf("Expected %s, got %s instead", taskName, l1[0].Task)
	}
	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Errorf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())
	if err := l1.Save(tf.Name()); err != nil {
		t.Errorf("Error saving list to file: %s", err)
	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Errorf("Error getting list from file: %s", err)
	}
	if l1[0].Task != l2[0].Task {
		t.Errorf("Expected %s, got %s instead", l1[0].Task, l2[0].Task)
	}
}
