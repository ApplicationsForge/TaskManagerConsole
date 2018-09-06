package task_terminal

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	assert := assert.New(t)
	app := &App{TaskTerminal: &TaskTerminal{}, TaskStore: &MemoryStore{}}
	year := strconv.Itoa(time.Now().Year())

	app.AddTask("a do some stuff due may 23")

	task := app.TaskTerminal.FindById(1)
	assert.Equal("do some stuff", task.Subject)
	assert.Equal(fmt.Sprintf("%s-05-23", year), task.Due)
	assert.Equal(false, task.Completed)
	assert.Equal(false, task.Archived)
	assert.Equal(false, task.IsPriority)
	assert.Equal("", task.CompletedDate)
	assert.Equal([]string{}, task.Tags)
	assert.Equal([]string{}, task.Users)
}

func TestAddDoneTask(t *testing.T) {
	assert := assert.New(t)
	app := &App{TaskTerminal: &TaskTerminal{}, TaskStore: &MemoryStore{}}

	app.AddDoneTask("Groked how to do done tasks @pop")

	task := app.TaskTerminal.FindById(1)
	assert.Equal("Groked how to do done tasks @pop", task.Subject)
	assert.Equal(true, task.Completed)
	assert.Equal(false, task.Archived)
	assert.Equal(false, task.IsPriority)
	assert.Equal([]string{}, task.Tags)
	assert.Equal(1, len(task.Users))
	assert.Equal("pop", task.Users[0])
}

func TestAddTaskWithEuropeanDates(t *testing.T) {
	assert := assert.New(t)
	app := &App{TaskTerminal: &TaskTerminal{}, TaskStore: &MemoryStore{}}
	year := strconv.Itoa(time.Now().Year())

	app.AddTask("a do some stuff due 23 may")

	task := app.TaskTerminal.FindById(1)
	assert.Equal("do some stuff", task.Subject)
	assert.Equal(fmt.Sprintf("%s-05-23", year), task.Due)
	assert.Equal(false, task.Completed)
	assert.Equal(false, task.Archived)
	assert.Equal(false, task.IsPriority)
	assert.Equal("", task.CompletedDate)
	assert.Equal([]string{}, task.Tags)
	assert.Equal([]string{}, task.Users)
}

func TestAddEmptyTask(t *testing.T) {
	assert := assert.New(t)
	app := &App{TaskTerminal: &TaskTerminal{}, TaskStore: &MemoryStore{}}

	app.AddTask("a")
	app.AddTask("a      ")
	app.AddTask("a\t\t\t\t")
	app.AddTask("a\t \t  \t   \t")

	assert.Equal(len(app.TaskTerminal.Data), 0)
}

func TestListbyTag(t *testing.T) {
	assert := assert.New(t)
	app := &App{TaskTerminal: &TaskTerminal{}, TaskStore: &MemoryStore{}}
	app.Load()

	// create three tasks w/wo a tag
	app.AddTask("this is a test +testme")
	app.AddTask("this is a test +testmetoo @work")
	app.AddTask("this is a test with no tags")
	app.CompleteTask("c 1")

	// simulate listTasks
	input := "l by t"
	filtered := NewFilter(app.TaskTerminal.Tasks()).Filter(input)
	grouped := app.getGroups(input, filtered)

	assert.Equal(3, len(grouped.Groups))

	// testme tag has 1 task and its completed
	assert.Equal(1, len(grouped.Groups["testme"]))
	assert.Equal(true, grouped.Groups["testme"][0].Completed)

	// testmetoo tag has 1 task and it has a user
	assert.Equal(1, len(grouped.Groups["testmetoo"]))
	assert.Equal(1, len(grouped.Groups["testmetoo"][0].Users))
	assert.Equal("work", grouped.Groups["testmetoo"][0].Users[0])
}

func TestListbyUser(t *testing.T) {
	assert := assert.New(t)
	app := &App{TaskTerminal: &TaskTerminal{}, TaskStore: &MemoryStore{}}
	app.Load()

	// create three tasks w/wo a user
	app.AddTask("this is a test +testme")
	app.AddTask("this is a test +testmetoo @work")
	app.AddTask("this is a test with no tags")
	app.CompleteTask("c 1")

	// simulate listTasks
	input := "l by u"
	filtered := NewFilter(app.TaskTerminal.Tasks()).Filter(input)
	grouped := app.getGroups(input, filtered)

	assert.Equal(2, len(grouped.Groups))

	// work user has 1 task and it has a tag of testmetoo
	assert.Equal(1, len(grouped.Groups["work"]))
	assert.Equal(1, len(grouped.Groups["work"][0].Tags))
	assert.Equal("testmetoo", grouped.Groups["work"][0].Tags[0])

	// There are two tasks with no user
	assert.Equal(2, len(grouped.Groups["No users"]))

	// check to see if the a tasks with no user contain a
	// completed task
	var hasACompletedTask bool
	for _, task := range grouped.Groups["No users"] {
		if task.Completed {
			hasACompletedTask = true
		}
	}
	assert.Equal(true, hasACompletedTask)
}

func TestGetId(t *testing.T) {
	assert := assert.New(t)
	app := &App{TaskTerminal: &TaskTerminal{}, TaskStore: &MemoryStore{}}
	// not a valid id
	assert.Equal(-1, app.getId("p"))
	// a single digit id
	assert.Equal(6, app.getId("6"))
	// a double digit id
	assert.Equal(66, app.getId("66"))
}

func TestGetIds(t *testing.T) {
	assert := assert.New(t)
	app := &App{TaskTerminal: &TaskTerminal{}, TaskStore: &MemoryStore{}}
	// no valid id here
	assert.Equal(0, len(app.getIds("p")))
	// one valid value here
	assert.Equal([]int{6}, app.getIds("6"))
	// lots of single post numbers
	assert.Equal([]int{6, 10, 8, 4}, app.getIds("6,10,8,4"))
	// a correct range
	assert.Equal([]int{6, 7, 8}, app.getIds("6-8"))
	// some incorrect ranges
	assert.Equal(0, len(app.getIds("6-6")))
	assert.Equal(0, len(app.getIds("8-6")))
	// some compsite ranges
	assert.Equal([]int{5, 6, 7, 8, 10, 11, 9}, app.getIds("5,6-8,10-11,9"))
}
