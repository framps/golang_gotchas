package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/framps/golang_gotchas/httpStress/dispatcher"
	"github.com/framps/golang_gotchas/httpStress/task.go"
	"github.com/framps/golang_gotchas/httpStress/utils"
	"github.com/framps/golang_gotchas/httpStress/worker"
)

const (
	workers = 3
	tasks   = 10
	debug   = true
)

func main() {

	if debug {
		utils.LogEnable()
	}

	utils.Logln("Starting http server ...")
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer svr.Close()

	dispatcher := dispatcher.NewDispatcher()

	utils.Log("Loading dispatcher with %d workers\n", workers)
	for i := 0; i < workers; i++ {
		dispatcher.WorkerAdd(worker.NewWorker(i))
	}

	utils.Log("Loading dispatcher with %d tasks\n", tasks)
	for i := 0; i < tasks; i++ {
		t := task.NewTask(svr.URL, "GET", true)
		dispatcher.TaskAdd(t)
	}

	utils.Log("Waiting for all workers to become ready for work\n")
	dispatcher.Trigger()
	fmt.Printf("Starting to work\n")

	dispatcher.Run()
	fmt.Printf("Waiting for work to complete\n")
	dispatcher.Wait()

	fmt.Printf("All work done\n")

	dispatcher.Statistics()
}
