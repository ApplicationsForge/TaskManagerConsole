package task_terminal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileStore(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "tasks.json"}
	tasks, _ := store.Load()
	assert.Equal(tasks[0].Subject, "this is the first subject", "")
}

func TestSave(t *testing.T) {
	store := &FileStore{FileLocation: "tasks.json"}
	tasks, _ := store.Load()
	store.Save(tasks)
}
