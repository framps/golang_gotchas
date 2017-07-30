package worker

import (
	"fmt"
	"net/http"
	"sync"

	task "github.com/framps/golang_gotchas/httpStress/task.go"
)

// Worker which processes a http work request
type Worker struct {
	ID           int
	Client       http.Client
	TaskChan     chan *task.Task
	FinishedWork int
}

// NewWorker -
func NewWorker(id int) *Worker {
	w := &Worker{ID: id}
	w.TaskChan = make(chan *task.Task)
	fmt.Printf("Created worker %d \n", w.ID)
	return w
}

// Run -
func (w *Worker) Run(workerChan chan *Worker, workerReadyWg *sync.WaitGroup) {
	workerReadyWg.Done()
	go func() {
		for {
			fmt.Printf("Worker %d: Ready for work\n", w.ID)
			workerChan <- w
			fmt.Printf("Worker %d: Waiting for work\n", w.ID)
			t := <-w.TaskChan
			fmt.Printf("Worker %d: Processing %v\n", w.ID, t)
			w.FinishedWork++
		}
	}()
}

// Statistics -
func (w *Worker) Statistics() string {
	return fmt.Sprintf("Worker %d: FinishedWork: %d", w.ID, w.FinishedWork)
}
