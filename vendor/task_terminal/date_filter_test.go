package task_terminal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFilterToday(t *testing.T) {
	assert := assert.New(t)

	var tasks []*Task
	todayTask := &Task{Id: 1, Subject: "one", Due: time.Now().Format("2006-01-02")}
	tomorrowTask := &Task{Id: 2, Subject: "two", Due: time.Now().AddDate(0, 0, 1).Format("2006-01-02")}
	tasks = append(tasks, todayTask)
	tasks = append(tasks, tomorrowTask)

	filter := NewDateFilter(tasks)
	filtered := filter.filterDueToday(time.Now())

	assert.Equal(1, len(filtered))
	assert.Equal(1, filtered[0].Id)
}

func TestFilterTomorrow(t *testing.T) {
	assert := assert.New(t)

	var tasks []*Task
	todayTask := &Task{Id: 1, Subject: "one", Due: time.Now().Format("2006-01-02")}
	tomorrowTask := &Task{Id: 2, Subject: "two", Due: time.Now().AddDate(0, 0, 1).Format("2006-01-02")}
	tasks = append(tasks, todayTask)
	tasks = append(tasks, tomorrowTask)

	filter := NewDateFilter(tasks)
	filtered := filter.filterDueTomorrow(time.Now())

	assert.Equal(1, len(filtered))
	assert.Equal(2, filtered[0].Id)
}

func TestFilterCompletedToday(t *testing.T) {
	assert := assert.New(t)

	var tasks []*Task
	taskNo1 := &Task{Id: 1, Subject: "one", Due: time.Now().Format("2006-01-02")}
	taskNo2 := &Task{Id: 2, Subject: "two", Due: time.Now().Format("2006-01-02")}

	tasks = append(tasks, taskNo1)
	tasks = append(tasks, taskNo2)

	filter := NewDateFilter(tasks)
	filtered := filter.filterCompletedToday(time.Now())

	assert.Equal(0, len(filtered))

	taskNo1.Complete()
	filtered = filter.filterCompletedToday(time.Now())

	assert.Equal(1, len(filtered))
	assert.Equal(1, filtered[0].Id)

	taskNo1.Uncomplete()
	taskNo2.Complete()
	filtered = filter.filterCompletedToday(time.Now())

	assert.Equal(1, len(filtered))
	assert.Equal(2, filtered[0].Id)

}

func TestFilterThisWeek(t *testing.T) {
	assert := assert.New(t)

	var tasks []*Task
	lastWeekTask := &Task{Id: 1, Subject: "two", Due: time.Now().AddDate(0, 0, -7).Format("2006-01-02")}
	todayTask := &Task{Id: 2, Subject: "one", Due: time.Now().Format("2006-01-02")}
	nextWeekTask := &Task{Id: 3, Subject: "two", Due: time.Now().AddDate(0, 0, 8).Format("2006-01-02")}
	tasks = append(tasks, lastWeekTask)
	tasks = append(tasks, todayTask)
	tasks = append(tasks, nextWeekTask)

	filter := NewDateFilter(tasks)
	filtered := filter.filterThisWeek(time.Now())

	assert.Equal(1, len(filtered))
	assert.Equal(2, filtered[0].Id)
}

func TestFilterCompletedThisWeek(t *testing.T) {
	assert := assert.New(t)

	var tasks []*Task
	lastWeekTask := &Task{Id: 1, Subject: "two", Due: time.Now().AddDate(0, 0, -7).Format("2006-01-02")}
	todayTask := &Task{Id: 2, Subject: "one", Due: time.Now().Format("2006-01-02")}
	nextWeekTask := &Task{Id: 3, Subject: "two", Due: time.Now().AddDate(0, 0, 8).Format("2006-01-02")}
	tasks = append(tasks, lastWeekTask)
	tasks = append(tasks, todayTask)
	tasks = append(tasks, nextWeekTask)

	filter := NewDateFilter(tasks)
	filtered := filter.filterCompletedThisWeek(time.Now())

	assert.Equal(0, len(filtered))

	todayTask.Complete()
	filtered = filter.filterCompletedThisWeek(time.Now())

	assert.Equal(1, len(filtered))
	assert.Equal(2, filtered[0].Id)

}

func TestFilterOverdue(t *testing.T) {
	assert := assert.New(t)

	var tasks []*Task
	lastWeekTask := &Task{Id: 1, Subject: "one", Due: time.Now().AddDate(0, 0, -7).Format("2006-01-02")}
	todayTask := &Task{Id: 2, Subject: "two", Due: bod(time.Now()).Format("2006-01-02")}
	tomorrowTask := &Task{Id: 3, Subject: "three", Due: time.Now().AddDate(0, 0, 1).Format("2006-01-02")}

	tasks = append(tasks, lastWeekTask)
	tasks = append(tasks, todayTask)
	tasks = append(tasks, tomorrowTask)

	filter := NewDateFilter(tasks)
	filtered := filter.filterOverdue(bod(time.Now()))

	assert.Equal(1, len(filtered))
	assert.Equal(1, filtered[0].Id)
}

func TestFilterDay(t *testing.T) {
	assert := assert.New(t)

	var tasks []*Task
	df := &DateFilter{}
	sunday := df.FindSunday(time.Now())

	mondayTask := &Task{Id: 1, Subject: "one", Due: sunday.AddDate(0, 0, 1).Format("2006-01-02")}
	tuesdayTask := &Task{Id: 2, Subject: "two", Due: sunday.AddDate(0, 0, 2).Format("2006-01-02")}

	tasks = append(tasks, mondayTask)
	tasks = append(tasks, tuesdayTask)

	filter := NewDateFilter(tasks)

	filtered := filter.filterDay(sunday, time.Monday)

	assert.Equal(1, len(filtered))
	assert.Equal(1, filtered[0].Id)
}

func TestFilterAgenda(t *testing.T) {
	assert := assert.New(t)

	var tasks []*Task

	completedTask := &Task{Id: 1, Subject: "completed", Completed: true, Due: time.Now().Format("2006-01-02")}
	uncompletedTask := &Task{Id: 2, Subject: "uncompleted", Due: time.Now().Format("2006-01-02")}

	tasks = append(tasks, completedTask)
	tasks = append(tasks, uncompletedTask)

	filter := NewDateFilter(tasks)

	filtered := filter.filterAgenda(time.Now())

	assert.Equal(1, len(filtered))
	assert.Equal(2, filtered[0].Id)
}
