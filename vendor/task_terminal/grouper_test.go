package task_terminal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupByUser(t *testing.T) {
	assert := assert.New(t)

	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)

	grouper := &Grouper{}
	grouped := grouper.GroupByUser(list.Tasks())

	assert.Equal(2, len(grouped.Groups["root"]), "")
	assert.Equal(1, len(grouped.Groups["more"]), "")
}

func TestGroupByTag(t *testing.T) {
	assert := assert.New(t)

	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)

	grouper := &Grouper{}
	grouped := grouper.GroupByTag(list.Tasks())

	assert.Equal(2, len(grouped.Groups["test1"]), "")
}
