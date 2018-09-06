package task_terminal

type MemoryStore struct {
	Tasks []*Task
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (m *MemoryStore) Initialize() {}

func (m *MemoryStore) Load() ([]*Task, error) {
	return m.Tasks, nil
}

func (m *MemoryStore) Save(tasks []*Task) {
	m.Tasks = tasks
}
