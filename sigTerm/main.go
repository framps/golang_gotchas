// Listen on sigterm and gracefully shutdown.
//
// See github.com/framps/golang_gotchas for latest code
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	goRoutines      = 100
	exitProbability = 0.5
	delay           = time.Second
	debug           = false
)

type counter struct {
	mutex   sync.Mutex
	counter int
}

var count counter

var wg sync.WaitGroup

func juggle(id int) {
	defer wg.Done()

	count.mutex.Lock()
	count.counter++
	count.mutex.Unlock()

	for {
		r := rand.Float32()

		if r < exitProbability {
			if debug {
				fmt.Printf("(%d) - Exiting: %d\n", count.counter, id)
			}
			count.mutex.Lock()
			count.counter--
			count.mutex.Unlock()
			return
		}
		if debug {
			fmt.Printf("(%d) - Sleeping: %d\n", count.counter, id)
		}
		time.Sleep(delay)
	}
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < goRoutines; i++ {
		wg.Add(1)
		go juggle(i)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		switch sig {
		case syscall.SIGTERM:
			fmt.Println("SIGTERM received")
		}
		fmt.Println("Awaiting SIGTERM")
		signal.Ignore(syscall.SIGINT, syscall.SIGTERM)
		done <- true
	}()

	go func(done chan bool) {
		fmt.Println("Waiting for all goroutines to terminate")
		wg.Wait()
		done <- true
	}(done)

	<-done
	fmt.Println("Graceful shutdown and waiting to goroutines to terminate")
	wg.Wait()
	fmt.Println("Exiting")

}