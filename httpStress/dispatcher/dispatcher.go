package dispatcher

import (
	"fmt"
	"sync"

	task "github.com/framps/golang_gotchas/httpStress/task.go"
	"github.com/framps/golang_gotchas/httpStress/worker"
)

// Dispatcher -
type Dispatcher struct {
	Tasks         map[int]*task.Task
	WorkerChan    chan *worker.Worker
	workerReadywg sync.WaitGroup
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
	fmt.Printf("Dispatcher: Adding task %d\n", task.ID)
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.Tasks[task.ID] = task
	d.Wg.Add(1)
}

// TaskRemove -
func (d *Dispatcher) TaskRemove(task task.Task) {
	fmt.Printf("Dispatcher: Removing task %d", task.ID)
	d.mutex.Lock()
	defer d.mutex.Unlock()
	delete(d.Tasks, task.ID)
	d.Wg.Done()
}

// WorkerAdd -
func (d *Dispatcher) WorkerAdd(worker *worker.Worker) {
	d.Workers = append(d.Workers, worker)
	d.workerReadywg.Add(1)
	worker.Run(d.WorkerChan, &d.workerReadywg)
}

// Run -
func (d *Dispatcher) Run() {
	fmt.Printf("Dispatcher: Running\n")
	for _, t := range d.Tasks {
		fmt.Printf("Dispatcher: Listening for free worker for task %d\n", t.ID)
		w := <-d.WorkerChan // wait for free worker
		fmt.Printf("Dispatcher: Found free worker %d\n", w.ID)
		w.TaskChan <- t // send task to worker
	}
}

// Trigger -
func (d *Dispatcher) Trigger() {
	fmt.Printf("Dispatcher: Waiting for worker to become ready\n")
	d.workerReadywg.Wait()
}

// Statistics -
func (d *Dispatcher) Statistics() {
	for _, w := range d.Workers {
		fmt.Printf("%s\n", w.Statistics())
	}
}
