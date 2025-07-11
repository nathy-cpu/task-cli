package tasks

import (
	"testing"
)

func TestAddTask(t *testing.T) {
	t.Run("Testing Adding new tasks", func(t *testing.T) {
		assertSuccess(t, AddTask(1, "Test task adding", TaskPending))
	})
}

func assertSuccess(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Adding task failed: %v\n", err)
	}
}
