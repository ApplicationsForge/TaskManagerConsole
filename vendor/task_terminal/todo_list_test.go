package task_terminal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextId(t *testing.T) {
	assert := assert.New(t)
	task := &Task{Subject: "testing", Completed: false, Archived: false}
	list := &TaskTerminal{}
	assert.Equal(1, list.NextId())
	list.Add(task)
	assert.Equal(2, list.NextId())
}

func TestNextIdWhenTaskDeleted(t *testing.T) {
	assert := assert.New(t)
	task := &Task{Subject: "testing", Completed: false, Archived: false}
	task2 := &Task{Subject: "testing2", Completed: false, Archived: false}
	task3 := &Task{Subject: "testing3", Completed: false, Archived: false}
	list := &TaskTerminal{}

	list.Add(task)
	list.Add(task2)
	list.Add(task3)

	list.Delete(2)
	assert.Equal(2, list.NextId())
	list.Add(task2)
	assert.Equal(4, list.NextId())
	list.Delete(1)
	assert.Equal(1, list.NextId())
}

func TestMaxId(t *testing.T) {
	assert := assert.New(t)
	task := &Task{Subject: "testing", Completed: false, Archived: false}
	task2 := &Task{Subject: "testing 2", Completed: false, Archived: false}
	list := &TaskTerminal{}
	assert.Equal(0, list.MaxId())
	list.Add(task)
	assert.Equal(1, list.MaxId())
	list.Add(task2)
	assert.Equal(2, list.MaxId())
}

func TestIndexOf(t *testing.T) {
	assert := assert.New(t)
	task := &Task{Subject: "Grant"}
	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)

	assert.Equal(-1, list.IndexOf(task))
	assert.Equal(0, list.IndexOf(list.Data[0]))
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)
	assert.Equal(2, len(list.Data))
	list.Delete(1)
	assert.Equal(1, len(list.Data))
}

func TestComplete(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)
	assert.Equal(false, list.FindById(1).Completed)
	list.Complete(1)
	assert.Equal(true, list.FindById(1).Completed)
}

func TestArchive(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)
	assert.Equal(false, list.FindById(2).Archived)
	list.Archive(2)
	assert.Equal(true, list.FindById(2).Archived)
}
func TestUnarchive(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)
	assert.Equal(true, list.FindById(1).Archived)
	list.Unarchive(1)
	assert.Equal(false, list.FindById(1).Archived)
}

func TestUncomplete(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)
	assert.Equal(true, list.FindById(2).Completed)
	list.Uncomplete(2)
	assert.Equal(false, list.FindById(2).Completed)
}

func TestGarbageCollect(t *testing.T) {
	assert := assert.New(t)
	list := &TaskTerminal{}
	task := &Task{Subject: "testing", Completed: false, Archived: true}
	task2 := &Task{Subject: "testing2", Completed: false, Archived: false}
	task3 := &Task{Subject: "testing3", Completed: false, Archived: true}
	list.Add(task)
	list.Add(task2)
	list.Add(task3)

	list.GarbageCollect()

	assert.Equal(len(list.Data), 1)
	assert.Equal(1, list.NextId())
	assert.Equal(2, list.MaxId())
}

func TestPrioritizeNotInTasksJson(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "tasks.json"}
	list := &TaskTerminal{}
	tasks, _ := store.Load()
	list.Load(tasks)
	assert.Equal(false, list.FindById(2).IsPriority)
}

func TestPrioritizeTask(t *testing.T) {
	assert := assert.New(t)
	list := &TaskTerminal{}
	task := &Task{Archived: false, Completed: false, Subject: "testing", IsPriority: false}
	list.Add(task)
	list.Prioritize(1)
	assert.Equal(true, list.FindById(1).IsPriority)
	list.Unprioritize(1)
	assert.Equal(false, list.FindById(1).IsPriority)
}
