package task_terminal

type Grouper struct{}

type GroupedTasks struct {
	Groups map[string][]*Task
}

func (g *Grouper) GroupByUser(tasks []*Task) *GroupedTasks {
	groups := map[string][]*Task{}

	allUsers := []string{}

	for _, task := range tasks {
		allUsers = AddIfNotThere(allUsers, task.Users)
	}

	for _, task := range tasks {
		for _, user := range task.Users {
			groups[user] = append(groups[user], task)
		}
		if len(task.Users) == 0 {
			groups["No users"] = append(groups["No users"], task)
		}
	}

	return &GroupedTasks{Groups: groups}
}

func (g *Grouper) GroupByTag(tasks []*Task) *GroupedTasks {
	groups := map[string][]*Task{}

	allTags := []string{}

	for _, task := range tasks {
		allTags = AddIfNotThere(allTags, task.Tags)
	}

	for _, task := range tasks {
		for _, tag := range task.Tags {
			groups[tag] = append(groups[tag], task)
		}
		if len(task.Tags) == 0 {
			groups["No Tags"] = append(groups["No Tags"], task)
		}
	}
	return &GroupedTasks{Groups: groups}
}

func (g *Grouper) GroupByNothing(tasks []*Task) *GroupedTasks {
	groups := map[string][]*Task{}
	groups["all"] = tasks
	return &GroupedTasks{Groups: groups}
}
