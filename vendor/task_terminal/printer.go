package task_terminal

type Printer interface {
	Print(*GroupedTasks, bool)
}
