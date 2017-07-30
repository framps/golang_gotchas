package main

// Copyright (C) 2017 framp at linux-tips-and-tricks dot de
//
// Sample how to kick off http requests as fast as possible
//
// See github.com/framps/golang_gotchas for latest code
//
// This code is based and was enhanced
// see http://stackoverflow.com/questions/23318419/how-can-i-effectively-max-out-concurrent-http-requests

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"runtime"
	"time"
)

var (
	reqs int
	max  int
)

func init() {
	flag.IntVar(&reqs, "reqs", 1000, "Total requests")
	flag.IntVar(&max, "concurrent", 200, "Maximum concurrent requests")
}

// Resp -
type Resp struct {
	*http.Response
	err error
}

func makeResponses(url string, reqs int, rc chan Resp, sem chan bool) {
	defer close(rc)
	defer close(sem)
	for reqs > 0 {
		select {
		case sem <- true:
			req, _ := http.NewRequest("GET", url, nil)
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

	flag.Parse()
	fmt.Printf("Starting local server ...\n")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//		fmt.Println("*")
	}))
	srv.Config.SetKeepAlivesEnabled(false)
	defer srv.Close()

	u, _ := url.Parse(srv.URL)

	fmt.Printf("Starting %d max concurrent requests of %d on %d procs ...\n", max, reqs, runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	rc := make(chan Resp)
	sem := make(chan bool, max)
	start := time.Now()
	go makeResponses(u.String(), reqs, rc, sem)
	conns := getResponses(rc)
	end := time.Since(start)
	fmt.Printf("Connections: %d\nTotal time: %s\n", conns, end)
}
