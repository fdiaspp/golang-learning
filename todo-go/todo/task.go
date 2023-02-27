package task

import (
	"time"
)

type Task struct {
	Name        string
	Description string
	DueDate     time.Time
}

func SetDone(t *Task) *Task {
	(*t).DueDate = time.Now()
	return t
}
