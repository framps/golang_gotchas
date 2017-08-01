package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/framps/golang_gotchas/httpStress/dispatcher"
	"github.com/framps/golang_gotchas/httpStress/task.go"
	"github.com/framps/golang_gotchas/httpStress/utils"
	"github.com/framps/golang_gotchas/httpStress/worker"
)

func main() {

	workers := flag.Int("w", 3, "Number of workers")
	tasks := flag.Int("t", 3, "Number of tasks")
	debug := flag.Bool("d", false, "Debug enabled")
	flag.Parse()

	if *debug {
		utils.LogEnable()
	}

	utils.Logln("Starting http server ...")
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte{byte('*')})
	}))
	defer svr.Close()

	dispatcher := dispatcher.NewDispatcher()

	utils.Log("Loading dispatcher with %d workers\n", workers)
	for i := 0; i < *workers; i++ {
		dispatcher.WorkerAdd(worker.NewWorker(i))
	}

	utils.Log("Loading dispatcher with %d tasks\n", tasks)
	for i := 0; i < *tasks; i++ {
		//t := task.NewTask(svr.URL, "GET", true)
		t := task.NewTask("http://www.linux-tips-and-tricks.de", "GET", true)
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
