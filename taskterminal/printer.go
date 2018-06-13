package taskterminal

type Printer interface {
	Print(*GroupedTasks, bool)
}
