package task_terminal

type Store interface {
	Initialize()
	Load() ([]*Task, error)
	Save(tasks []*Task)
}
