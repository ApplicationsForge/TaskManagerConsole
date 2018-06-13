package taskterminal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterArchived(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)
	filter := NewFilter(list.Tasks())
	archived := filter.filterArchived("l archived")
	assert.Equal(1, len(archived))
	assert.Equal(true, archived[0].Archived)
}

func TestFilterUnarchivedByDefault(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)
	filter := NewFilter(list.Tasks())
	unarchived := filter.filterArchived("l")
	assert.Equal(1, len(unarchived))
	assert.Equal(false, unarchived[0].Archived)
}

func TestFilterShowArchivedWhenWeAskForCompleted(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)
	filter := NewFilter(list.Tasks())
	unarchived := filter.filterArchived("completed")
	assert.Equal(2, len(unarchived))
	assert.Equal(false, unarchived[0].Archived)
	assert.Equal(true, unarchived[1].Archived)
}

func TestGetArchived(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)
	filter := NewFilter(list.Tasks())
	archived := filter.getArchived()
	assert.Equal(1, len(archived))
	assert.Equal(true, archived[0].Archived)
}

func TestGetUnarchived(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)
	filter := NewFilter(list.Tasks())
	unarchived := filter.getUnarchived()
	assert.Equal(1, len(unarchived))
	assert.Equal(false, unarchived[0].Archived)
}
