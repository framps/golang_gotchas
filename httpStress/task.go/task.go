package task

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var taskID int

// WorkerTask -
type WorkerTask interface {
	Receive(client *http.Client) int // receive http request
	PostProcess()                    // process received request if sync request
	IsSyncRequest() bool
}

// SimpleWorkerTask - Just a simple http get
type SimpleWorkerTask struct {
	Task
}

// Task - describes the task a worker has to execute
type Task struct {
	ID       int
	URL      string
	Request  string // GET, POST, ...
	wait     bool   // whether to wait for response
	Response []byte // response
}

// NewTask -
func NewTask(url, request string, wait bool) *Task {
	taskID++
	return &Task{ID: taskID, URL: url, Request: request, wait: wait}
}

func (t *Task) String() string {
	return fmt.Sprintf("Task %d: %s %s", t.ID, t.Request, t.URL)
}

// Receive -
func (t *Task) Receive(client *http.Client) int {
	rsp, err := client.Get(t.URL)
	if err != nil {
		panic(err)
	}
	t.Response, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	rsp.Body.Close()
	return rsp.StatusCode
}

// PostProcess -
func (t *Task) PostProcess() {
	// fmt.Printf("Response: %s\n", t.Response[:32])
}

// IsSyncRequest -
func (t *Task) IsSyncRequest() bool {
	return t.wait
}
