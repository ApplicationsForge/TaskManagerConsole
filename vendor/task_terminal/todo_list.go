package task_terminal

import "sort"

type TaskTerminal struct {
	Data []*Task
}

func (t *TaskTerminal) Load(tasks []*Task) {
	t.Data = tasks
}

func (t *TaskTerminal) Add(task *Task) {
	task.Id = t.NextId()
	t.Data = append(t.Data, task)
}

func (t *TaskTerminal) Delete(ids ...int) {
	for _, id := range ids {
		task := t.FindById(id)
		if task == nil {
			continue
		}
		i := -1
		for index, task := range t.Data {
			if task.Id == id {
				i = index
			}
		}

		t.Data = append(t.Data[:i], t.Data[i+1:]...)
	}
}

func (t *TaskTerminal) ChangeTaskStatus(status string, ids ...int) {
	for _, id := range ids {
		task := t.FindById(id)
		if task == nil {
			continue
		}
		task.ChangeStatus(status)
		t.Delete(id)
		t.Data = append(t.Data, task)
	}
}

func (t *TaskTerminal) Archive(ids ...int) {
	for _, id := range ids {
		task := t.FindById(id)
		if task == nil {
			continue
		}
		task.Archive()
		t.Delete(id)
		t.Data = append(t.Data, task)
	}
}

func (t *TaskTerminal) Unarchive(ids ...int) {
	for _, id := range ids {
		task := t.FindById(id)
		if task == nil {
			continue
		}
		task.Unarchive()
		t.Delete(id)
		t.Data = append(t.Data, task)
	}
}

func (t *TaskTerminal) Prioritize(ids ...int) {
	for _, id := range ids {
		task := t.FindById(id)
		if task == nil {
			continue
		}
		task.Prioritize()
		t.Delete(id)
		t.Data = append(t.Data, task)
	}
}

func (t *TaskTerminal) Unprioritize(ids ...int) {
	for _, id := range ids {
		task := t.FindById(id)
		if task == nil {
			continue
		}
		task.Unprioritize()
		t.Delete(id)
		t.Data = append(t.Data, task)
	}
}

func (t *TaskTerminal) IndexOf(taskToFind *Task) int {
	for i, task := range t.Data {
		if task.Id == taskToFind.Id {
			return i
		}
	}
	return -1
}

type ByDate []*Task

func (a ByDate) Len() int      { return len(a) }
func (a ByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool {
	t1Due := a[i].CalculateDueTime()
	t2Due := a[j].CalculateDueTime()
	return t1Due.Before(t2Due)
}

func (t *TaskTerminal) Tasks() []*Task {
	sort.Sort(ByDate(t.Data))
	return t.Data
}

func (t *TaskTerminal) MaxId() int {
	maxId := 0
	for _, task := range t.Data {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	return maxId
}

func (t *TaskTerminal) NextId() int {
	var found bool
	maxID := t.MaxId()
	for i := 1; i <= maxID; i++ {
		found = false
		for _, task := range t.Data {
			if task.Id == i {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}
	return maxID + 1
}

func (t *TaskTerminal) FindById(id int) *Task {
	for _, task := range t.Data {
		if task.Id == id {
			return task
		}
	}
	return nil
}

func (t *TaskTerminal) GarbageCollect() {
	var toDelete []*Task
	for _, task := range t.Data {
		if task.Archived {
			toDelete = append(toDelete, task)
		}
	}
	for _, task := range toDelete {
		t.Delete(task.Id)
	}
}
