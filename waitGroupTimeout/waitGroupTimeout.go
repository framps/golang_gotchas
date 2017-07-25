// Waitgroup which times out
//
// See github.com/framps/golang_gotchas for latest code
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var flip bool
	for {

		var wg sync.WaitGroup
		wg.Add(1)

		timeout := time.Second * 5
		fmt.Printf("Waitgroup timeout: %s\n", timeout.String())

		go func() {
			var to time.Duration
			if flip {
				to = time.Second * 3 // force timeout
			} else {
				to = timeout * 3 // force wait to succeed
			}
			flip = !flip
			fmt.Printf("Go func waiting: %s\n", to.String())
			time.Sleep(to)
			wg.Done()
			fmt.Println("Go func done")
		}()

		if waitTimeout(&wg, timeout) {
			fmt.Println("Waitgroup timed out")
		} else {
			fmt.Println("Wait group finished")
		}
		fmt.Println("Waiting to start new loop")
		time.Sleep(3 * time.Second)
	}
}

// waitTimeout waits for the waitgroup for the specified timeout
// Returns true if timed out
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // wait succeeded
	case <-time.After(timeout):
		return true // wait timed out
	}
}
