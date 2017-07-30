package main

// see http://stackoverflow.com/questions/23318419/how-can-i-effectively-max-out-concurrent-http-requests

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

// Resph -
type Resp struct {
	*http.Response
	err error
}

func makeResponses(reqs int, rc chan Resp, sem chan bool) {
	defer close(rc)
	defer close(sem)
	for reqs > 0 {
		select {
		case sem <- true:
			req, _ := http.NewRequest("GET", "http://localhost:/", nil)
			transport := &http.Transport{}
			resp, err := transport.RoundTrip(req)
			r := Resp{resp, err}
			rc <- r
			reqs--
		default:
			<-sem
		}
	}
}

func getResponses(rc chan Resp) int {
	conns := 0
	for {
		select {
		case r, ok := <-rc:
			if ok {
				conns++
				if r.err != nil {
					fmt.Println(r.err)
				} else {
					// Do something with response
					if err := r.Body.Close(); err != nil {
						fmt.Println(r.err)
					}
				}
			} else {
				return conns
			}
		}
	}
}

func main() {
	// reqs := 10000
	reqs := 10
	// maxConcurrent := 1000
	maxConcurrent := 10
	runtime.GOMAXPROCS(runtime.NumCPU())
	rc := make(chan Resp)
	sem := make(chan bool, maxConcurrent)
	start := time.Now()
	go makeResponses(reqs, rc, sem)
	conns := getResponses(rc)
	end := time.Since(start)
	fmt.Printf("Connections: %d\nTotal time: %s\n", conns, end)
}
