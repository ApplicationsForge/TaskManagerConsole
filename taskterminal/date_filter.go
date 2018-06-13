package taskterminal

import (
	"regexp"
	"time"
)

type DateFilter struct {
	Tasks    []*Task
	Location *time.Location
}

func NewDateFilter(tasks []*Task) *DateFilter {
	return &DateFilter{Tasks: tasks, Location: time.Now().Location()}
}

func filterOnDue(task *Task) string {
	return task.Due
}

func filterOnCompletedDate(task *Task) string {
	return task.CompletedDateToDate()
}

func (f *DateFilter) FilterDate(input string) []*Task {
	agendaRegex, _ := regexp.Compile(`agenda.*$`)
	if agendaRegex.MatchString(input) {
		return f.filterAgenda(bod(time.Now()))
	}

	// filter due items
	r, _ := regexp.Compile(`due .*$`)
	match := r.FindString(input)
	switch {
	case match == "due tod" || match == "due today":
		return f.filterDueToday(bod(time.Now()))
	case match == "due tom" || match == "due tomorrow":
		return f.filterDueTomorrow(bod(time.Now()))
	case match == "due sun" || match == "due sunday":
		return f.filterDay(bod(time.Now()), time.Sunday)
	case match == "due mon" || match == "due monday":
		return f.filterDay(bod(time.Now()), time.Monday)
	case match == "due tue" || match == "due tuesday":
		return f.filterDay(bod(time.Now()), time.Tuesday)
	case match == "due wed" || match == "due wednesday":
		return f.filterDay(bod(time.Now()), time.Wednesday)
	case match == "due thu" || match == "due thursday":
		return f.filterDay(bod(time.Now()), time.Thursday)
	case match == "due fri" || match == "due friday":
		return f.filterDay(bod(time.Now()), time.Friday)
	case match == "due sat" || match == "due saturday":
		return f.filterDay(bod(time.Now()), time.Saturday)
	case match == "due this week":
		return f.filterThisWeek(bod(time.Now()))
	case match == "due next week":
		return f.filterNextWeek(bod(time.Now()))
	case match == "due last week":
		return f.filterLastWeek(bod(time.Now()))
	case match == "overdue":
		return f.filterOverdue(bod(time.Now()))
	}

	// filter completed items
	r, _ = regexp.Compile(`completed .*$`)
	match = r.FindString(input)
	switch {
	case match == "completed tod" || match == "completed today":
		return f.filterCompletedToday(bod(time.Now()))
	case match == "completed this week":
		return f.filterCompletedThisWeek(bod(time.Now()))
	}

	return f.Tasks
}

func (f *DateFilter) filterAgenda(pivot time.Time) []*Task {
	var ret []*Task

	for _, task := range f.Tasks {
		if task.Due == "" || task.Status != "Task" {
			continue
		}
		dueTime, _ := time.ParseInLocation("2006-01-02", task.Due, f.Location)
		if dueTime.Before(pivot) || task.Due == pivot.Format("2006-01-02") {
			ret = append(ret, task)
		}
	}
	return ret
}

func (f *DateFilter) filterToExactDate(pivot time.Time, filterOn func(*Task) string) []*Task {
	var ret []*Task
	for _, task := range f.Tasks {
		if filterOn(task) == pivot.Format("2006-01-02") {
			ret = append(ret, task)
		}
	}
	return ret
}

func (f *DateFilter) filterDueToday(pivot time.Time) []*Task {
	return f.filterToExactDate(pivot, filterOnDue)
}

func (f *DateFilter) filterDueTomorrow(pivot time.Time) []*Task {
	pivot = pivot.AddDate(0, 0, 1)
	return f.filterToExactDate(pivot, filterOnDue)
}

func (f *DateFilter) filterCompletedToday(pivot time.Time) []*Task {
	return f.filterToExactDate(pivot, filterOnCompletedDate)
}

func (f *DateFilter) filterDay(pivot time.Time, day time.Weekday) []*Task {
	thisWeek := NewDateFilter(f.filterThisWeek(pivot))
	pivot = f.FindSunday(pivot).AddDate(0, 0, int(day))
	return thisWeek.filterToExactDate(pivot, filterOnDue)
}

func (f *DateFilter) filterBetweenDatesInclusive(begin, end time.Time, filterOn func(*Task) string) []*Task {
	var ret []*Task

	for _, task := range f.Tasks {
		dueTime, _ := time.ParseInLocation("2006-01-02", filterOn(task), f.Location)
		if (begin.Before(dueTime) || begin.Equal(dueTime)) && end.After(dueTime) {
			ret = append(ret, task)
		}
	}
	return ret
}

func (f *DateFilter) filterThisWeek(pivot time.Time) []*Task {

	begin := bod(f.FindSunday(pivot))
	end := begin.AddDate(0, 0, 7)

	return f.filterBetweenDatesInclusive(begin, end, filterOnDue)
}

func (f *DateFilter) filterCompletedThisWeek(pivot time.Time) []*Task {

	begin := bod(f.FindSunday(pivot))
	end := begin.AddDate(0, 0, 7)

	return f.filterBetweenDatesInclusive(begin, end, filterOnCompletedDate)
}

func (f *DateFilter) filterBetweenDatesExclusive(begin, end time.Time) []*Task {
	var ret []*Task

	for _, task := range f.Tasks {
		dueTime, _ := time.ParseInLocation("2006-01-02", task.Due, f.Location)
		if begin.Before(dueTime) && end.After(dueTime) {
			ret = append(ret, task)
		}
	}
	return ret
}

func (f *DateFilter) filterNextWeek(pivot time.Time) []*Task {

	begin := f.FindSunday(pivot).AddDate(0, 0, 7)
	end := begin.AddDate(0, 0, 7)

	return f.filterBetweenDatesExclusive(begin, end)
}

func (f *DateFilter) filterLastWeek(pivot time.Time) []*Task {

	begin := f.FindSunday(pivot).AddDate(0, 0, -7)
	end := begin.AddDate(0, 0, 7)

	return f.filterBetweenDatesExclusive(begin, end)
}

func (f *DateFilter) filterOverdue(pivot time.Time) []*Task {
	var ret []*Task

	pivotDate := pivot.Format("2006-01-02")

	for _, task := range f.Tasks {
		dueTime, _ := time.ParseInLocation("2006-01-02", task.Due, f.Location)
		if dueTime.Before(pivot) && pivotDate != task.Due {
			ret = append(ret, task)
		}
	}
	return ret
}

func (f *DateFilter) FindSunday(pivot time.Time) time.Time {
	switch pivot.Weekday() {
	case time.Sunday:
		return pivot
	case time.Monday:
		return pivot.AddDate(0, 0, -1)
	case time.Tuesday:
		return pivot.AddDate(0, 0, -2)
	case time.Wednesday:
		return pivot.AddDate(0, 0, -3)
	case time.Thursday:
		return pivot.AddDate(0, 0, -4)
	case time.Friday:
		return pivot.AddDate(0, 0, -5)
	case time.Saturday:
		return pivot.AddDate(0, 0, -6)
	}
	return pivot
}
