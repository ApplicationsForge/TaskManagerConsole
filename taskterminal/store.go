package taskterminal

type Store interface {
	Initialize()
	Load() ([]*Task, error)
	Save(tasks []*Task)
}
