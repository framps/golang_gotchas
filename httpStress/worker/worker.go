package worker

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	task "github.com/framps/golang_gotchas/httpStress/task.go"
	"github.com/framps/golang_gotchas/httpStress/utils"
)

// Statistics -

// Statistics -
type Statistics struct {
	SumDuration time.Duration
	Requests    int
	rc200       int
	rc300       int
	rc400       int
	rc500       int
}

// Worker which processes a http work request
type Worker struct {
	ID        int
	Client    *http.Client
	TaskChan  chan task.WorkerTask
	Statistic Statistics
}

// NewWorker -
func NewWorker(id int) *Worker {
	w := &Worker{ID: id}
	w.TaskChan = make(chan task.WorkerTask)

	w.Client = &http.Client{}
	utils.Log("Created worker %d \n", w.ID)
	return w
}

// Run -
func (w *Worker) Run(workerChan chan *Worker, workerReadyWg *sync.WaitGroup, workerBusyWg *sync.WaitGroup) {
	workerReadyWg.Done()
	go func() {
		for {
			utils.Log("Worker %d: Ready for work\n", w.ID)
			workerChan <- w
			utils.Log("Worker %d: Waiting for work\n", w.ID)
			workerBusyWg.Add(1)
			t := <-w.TaskChan
			utils.Log("Worker %d: Processing %v\n", w.ID, t)
			requestStart := time.Now()
			statusCode := t.Receive(w.Client)
			if t.IsSyncRequest() {
				t.PostProcess()
				switch {
				case statusCode >= 200 && statusCode < 300:
					w.Statistic.rc200++
				case statusCode >= 300 && statusCode < 400:
					w.Statistic.rc300++
				case statusCode >= 400 && statusCode < 400:
					w.Statistic.rc400++
				case statusCode >= 500 && statusCode < 500:
					w.Statistic.rc500++
				}
			}
			requestDuration := time.Since(requestStart)
			w.Statistic.SumDuration += requestDuration
			w.Statistic.Requests++
			workerBusyWg.Done()
		}
	}()
}

// Statistics -
func (w *Worker) Statistics() string {
	d := float64(0)
	if w.Statistic.Requests > 0 {
		d = w.Statistic.SumDuration.Seconds() / float64(w.Statistic.Requests)
	}

	return fmt.Sprintf("Worker %3d: n: %5d (#) duration avg: %3.3f (s)\n2xx: %d - 3xx: %d - 4xx: %d - 5xx: %d",
		w.ID,
		w.Statistic.Requests,
		d,
		w.Statistic.rc200,
		w.Statistic.rc300,
		w.Statistic.rc400,
		w.Statistic.rc500,
	)
}
