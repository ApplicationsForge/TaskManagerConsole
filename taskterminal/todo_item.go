package taskterminal

import "time"

// Timestamp format to include date, time with timezone support. Easy to parse
const ISO8601_TIMESTAMP_FORMAT = "2006-01-02T15:04:05Z07:00"

type Task struct {
	Id            int      `json:"id"`
	Title         string `json:"title"`
	Subject       string   `json:"subject"`
	Tags          []string `json:"tags"`
	Users         []string `json:"users"`
	Due           string   `json:"due"`
	Status        string   `json:"status"`
	CompletedDate string   `json:"completedDate"`
	Archived      bool     `json:"archived"`
	IsPriority    bool     `json:"isPriority"`
	Notes         []string `json:"notes"`
}

func NewTask() *Task {
	return &Task{Status: "ToDo", Archived: false, IsPriority: false}
}

func (t Task) Valid() bool {
	return (t.Subject != "")
}

func (t Task) CalculateDueTime() time.Time {
	if t.Due != "" {
		parsedTime, _ := time.Parse("2006-01-02", t.Due)
		return parsedTime
	} else {
		parsedTime, _ := time.Parse("2006-01-02", "1900-01-01")
		return parsedTime
	}
}

func (t *Task) ChangeStatus(status string) {
	if len(status) > 0 {
		t.Status = status
		t.CompletedDate = timestamp(time.Now()).Format(ISO8601_TIMESTAMP_FORMAT)
	}
}

func (t *Task) Archive() {
	t.Archived = true
}

func (t *Task) Unarchive() {
	t.Archived = false
}

func (t *Task) Prioritize() {
	t.IsPriority = true
}

func (t *Task) Unprioritize() {
	t.IsPriority = false
}

func (t Task) CompletedDateToDate() string {
	parsedTime, _ := time.Parse(ISO8601_TIMESTAMP_FORMAT, t.CompletedDate)
	return parsedTime.Format("2006-01-02")
}
