package task_terminal

import "regexp"

type TaskFilter struct {
	Tasks []*Task
}

func NewFilter(tasks []*Task) *TaskFilter {
	return &TaskFilter{Tasks: tasks}
}

func (f *TaskFilter) Filter(input string) []*Task {
	f.Tasks = f.filterArchived(input)
	f.Tasks = f.filterPrioritized(input)
	f.Tasks = f.filterTags(input)
	f.Tasks = f.filterUsers(input)
	f.Tasks = NewDateFilter(f.Tasks).FilterDate(input)

	return f.Tasks
}

func (t *TaskFilter) isFilteringByTags(input string) bool {
	parser := &Parser{}
	return len(parser.Tags(input)) > 0
}

func (t *TaskFilter) isFilteringByUsers(input string) bool {
	parser := &Parser{}
	return len(parser.Users(input)) > 0
}

func (f *TaskFilter) filterArchived(input string) []*Task {

	// do not filter archived if want completed items
	completedRegex, _ := regexp.Compile(`completed`)
	if completedRegex.MatchString(input) {
		return f.Tasks
	}

	r, _ := regexp.Compile(`ln? archived$`)
	if r.MatchString(input) {
		return f.getArchived()
	} else {
		return f.getUnarchived()
	}
}

func (f *TaskFilter) filterPrioritized(input string) []*Task {
	r, _ := regexp.Compile(`ln? p`)
	if r.MatchString(input) {
		return f.getPrioritized()
	} else {
		return f.Tasks
	}
}

func (f *TaskFilter) filterTags(input string) []*Task {
	if !f.isFilteringByTags(input) {
		return f.Tasks
	}
	parser := &Parser{}
	tags := parser.Tags(input)
	var ret []*Task

	for _, task := range f.Tasks {
		for _, taskTag := range task.Tags {
			for _, tag := range tags {
				if tag == taskTag {
					ret = AddTaskIfNotThere(ret, task)
				}
			}
		}
	}
	return ret
}

func (f *TaskFilter) filterUsers(input string) []*Task {
	if !f.isFilteringByUsers(input) {
		return f.Tasks
	}
	parser := &Parser{}
	users := parser.Users(input)
	var ret []*Task

	for _, task := range f.Tasks {
		for _, taskUser := range task.Users {
			for _, user := range users {
				if user == taskUser {
					ret = AddTaskIfNotThere(ret, task)
				}
			}
		}
	}
	return ret
}

func (f *TaskFilter) getArchived() []*Task {
	var ret []*Task
	for _, task := range f.Tasks {
		if task.Archived {
			ret = append(ret, task)
		}
	}
	return ret
}

func (f *TaskFilter) getPrioritized() []*Task {
	var ret []*Task
	for _, task := range f.Tasks {
		if task.IsPriority {
			ret = append(ret, task)
		}
	}
	return ret
}

func (f *TaskFilter) getUnarchived() []*Task {
	var ret []*Task
	for _, task := range f.Tasks {
		if !task.Archived {
			ret = append(ret, task)
		}
	}
	return ret
}
