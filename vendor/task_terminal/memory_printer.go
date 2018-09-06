package task_terminal

type MemoryPrinter struct {
	Groups *GroupedTasks
}

func (m *MemoryPrinter) Print(groupedTasks *GroupedTasks, printNotes bool) {
	m.Groups = groupedTasks
}
