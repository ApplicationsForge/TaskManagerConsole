package taskterminal

import "testing"

func TestNewTask(t *testing.T) {
	task := NewTask()

	if task.Completed || task.Archived || task.CompletedDate != "" {
		t.Error("Completed should be false for new tasks")
	}
}

func TestValidity(t *testing.T) {
	task := &Task{Subject: "test"}
	if !task.Valid() {
		t.Error("Expected valid task to be valid")
	}

	invalidTask := &Task{Subject: ""}
	if invalidTask.Valid() {
		t.Error("Invalid task is being reported as valid")
	}
}
