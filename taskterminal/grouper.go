package taskterminal

type Grouper struct{}

type GroupedTasks struct {
	Groups map[string][]*Task
}

func (g *Grouper) GroupByContext(tasks []*Task) *GroupedTasks {
	groups := map[string][]*Task{}

	allContexts := []string{}

	for _, task := range tasks {
		allContexts = AddIfNotThere(allContexts, task.Contexts)
	}

	for _, task := range tasks {
		for _, context := range task.Contexts {
			groups[context] = append(groups[context], task)
		}
		if len(task.Contexts) == 0 {
			groups["No contexts"] = append(groups["No contexts"], task)
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
