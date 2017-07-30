package main

import (
	"fmt"

	"github.com/framps/golang_gotchas/httpStress/dispatcher"
	"github.com/framps/golang_gotchas/httpStress/task.go"
	"github.com/framps/golang_gotchas/httpStress/worker"
)

const (
	workers = 20
	tasks   = 100
	debug   = true
)

func main() {

	dispatcher := dispatcher.NewDispatcher()

	fmt.Printf("Loading dispatcher with %d workers\n", workers)
	for i := 0; i < workers; i++ {
		dispatcher.WorkerAdd(worker.NewWorker(i))
	}

	fmt.Printf("Loading dispatcher with %d tasks\n", tasks)
	for i := 0; i < tasks; i++ {
		t := task.NewTask("http://www.google.de", "GET", true)
		dispatcher.TaskAdd(t)
	}

	fmt.Printf("Waiting for all workers to become ready for work\n")
	dispatcher.Trigger()
	fmt.Printf("Starting to work\n")
	dispatcher.Run()

	fmt.Printf("All work done\n")

	dispatcher.Statistics()
}
