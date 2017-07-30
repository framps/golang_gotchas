package dispatcher

import (
	"fmt"
	"sync"

	task "github.com/framps/golang_gotchas/httpStress/task.go"
	"github.com/framps/golang_gotchas/httpStress/utils"
	"github.com/framps/golang_gotchas/httpStress/worker"
)

// Dispatcher -
type Dispatcher struct {
	Tasks         map[int]*task.Task
	WorkerChan    chan *worker.Worker
	workerReadywg sync.WaitGroup
	workerBusyWg  sync.WaitGroup
	Workers       []*worker.Worker
	Wg            sync.WaitGroup
	mutex         sync.Mutex
}

// NewDispatcher -
func NewDispatcher() *Dispatcher {
	d := &Dispatcher{}
	d.Tasks = make(map[int]*task.Task)
	d.Workers = make([]*worker.Worker, 0)
	d.WorkerChan = make(chan *worker.Worker)
	return d
}

// TaskAdd -
func (d *Dispatcher) TaskAdd(task *task.Task) {
	utils.Log("Dispatcher: Adding task %d\n", task.ID)
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.Tasks[task.ID] = task
	d.Wg.Add(1)
}

// WorkerAdd -
func (d *Dispatcher) WorkerAdd(worker *worker.Worker) {
	d.Workers = append(d.Workers, worker)
	d.workerReadywg.Add(1)
	worker.Run(d.WorkerChan, &d.workerReadywg, &d.workerBusyWg)
}

// Wait -
func (d *Dispatcher) Wait() {
	d.workerBusyWg.Wait()
}

// Run -
func (d *Dispatcher) Run() {
	utils.Log("Dispatcher: Running\n")
	for _, t := range d.Tasks {
		utils.Log("Dispatcher: Listening for free worker for task %d\n", t.ID)
		w := <-d.WorkerChan // wait for free worker
		utils.Log("Dispatcher: Found free worker %d\n", w.ID)
		w.TaskChan <- t // send task to worker
	}
}

// Trigger -
func (d *Dispatcher) Trigger() {
	utils.Log("Dispatcher: Waiting for worker to become ready\n")
	d.workerReadywg.Wait()
}

// Statistics -
func (d *Dispatcher) Statistics() {
	for _, w := range d.Workers {
		fmt.Printf("%s\n", w.Statistics())
	}
}
