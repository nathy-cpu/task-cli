package tasks_test

import (
	"testing"

	tasks "github.com/nathy-cpu/task-cli"
)

func TestAddTask(t *testing.T) {
	list := tasks.List{}

	task := "New todo task"
	list.Add(task)

	if task != list[0].Description {
		t.Errorf("want: %q got: %q", task, list[0].Description)
	}
}

func TestCompleteTask(t *testing.T) {
	list := tasks.List{}

	task := "New todo task"
	list.Add(task)

	if task != list[0].Description {
		t.Errorf("want: %q got: %q", task, list[0].Description)
	}

	if list[0].Done {
		t.Errorf("New task should not be complete!")
	}

	list.Complete(1)
	if !list[0].Done {
		t.Errorf("New task should be complete!")
	}
}

func TestDeleteTask(t *testing.T) {
	list := tasks.List{}

	taskItems := []string{"New task", "Newer task", "Newest task"}

	for i, item := range taskItems {
		list.Add(item)
		if item != list[i].Description {
			t.Errorf("want: %q got: %q", item, list[i].Description)
		}
	}

	list.Delete(len(taskItems) / 2)

	if len(list) != len(taskItems) - 1 {
		t.Errorf("expected list length: %d, got: %d", len(taskItems) - 1, len(list))
	}

	if list[len(list) - 1].Description != taskItems[len(taskItems) - 1] {
		t.Errorf("expected list length: %q, got: %q", list[len(list) - 1].Description, taskItems[len(taskItems) - 1])
	}

}

