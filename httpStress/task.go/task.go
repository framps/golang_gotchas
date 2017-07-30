package task

import "fmt"

var taskID int

// Task - describes the task a worker has to execute
type Task struct {
	ID       int
	URL      string
	Request  string  // GET, POST, ...
	Wait     bool    // whether to wait for response
	Response *[]byte // response
}

// NewTask -
func NewTask(url, request string, wait bool) *Task {
	taskID++
	return &Task{ID: taskID, URL: url, Request: request, Wait: wait}
}

func (t *Task) String() string {
	return fmt.Sprintf("Task %d: %s %s", t.ID, t.Request, t.URL)
}
