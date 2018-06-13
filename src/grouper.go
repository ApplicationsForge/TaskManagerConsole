package todolist

type Grouper struct{}

type GroupedTodos struct {
	Groups map[string][]*Todo
}

func (g *Grouper) GroupByContext(todos []*Todo) *GroupedTodos {
	groups := map[string][]*Todo{}

	allContexts := []string{}

	for _, todo := range todos {
		allContexts = AddIfNotThere(allContexts, todo.Contexts)
	}

	for _, todo := range todos {
		for _, context := range todo.Contexts {
			groups[context] = append(groups[context], todo)
		}
		if len(todo.Contexts) == 0 {
			groups["No contexts"] = append(groups["No contexts"], todo)
		}
	}

	return &GroupedTodos{Groups: groups}
}

func (g *Grouper) GroupByTag(todos []*Todo) *GroupedTodos {
	groups := map[string][]*Todo{}

	allTags := []string{}

	for _, todo := range todos {
		allTags = AddIfNotThere(allTags, todo.Tags)
	}

	for _, todo := range todos {
		for _, tag := range todo.Tags {
			groups[tag] = append(groups[tag], todo)
		}
		if len(todo.Tags) == 0 {
			groups["No Tags"] = append(groups["No Tags"], todo)
		}
	}
	return &GroupedTodos{Groups: groups}
}

func (g *Grouper) GroupByNothing(todos []*Todo) *GroupedTodos {
	groups := map[string][]*Todo{}
	groups["all"] = todos
	return &GroupedTodos{Groups: groups}
}
