package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	threads = 10
	sleep   = time.Millisecond
	loops   = 3
)

func main() {
	start := time.Now()
	ch := make(chan string)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(os.Args[1:]))
	for _, url := range os.Args[1:] {
		go fetch(&waitGroup, url, ch, threads, sleep, loops)
	}

	fmt.Println("Waiting ...")
	waitGroup.Wait()

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(waitgroup *sync.WaitGroup, url string, ch chan<- string, threads int, sleep time.Duration, loops int) {

	fmt.Print("fetching\n")
	var wg sync.WaitGroup
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		go func() {
			fmt.Printf("Starting %s\n", url)
			for i := 0; i < loops; i++ {
				start := time.Now()
				resp, err := http.Get(url)
				if err != nil {
					fmt.Print(err)
					return
				}
				nbytes, err := io.Copy(ioutil.Discard, resp.Body)
				resp.Body.Close()
				if err != nil {
					fmt.Print("While reading %s: %v", url, err)
					return
				}
				secs := time.Since(start).Seconds()
				fmt.Printf("%.2fs  %7d  %s\n", secs, nbytes, url)

				time.Sleep(sleep)
			}
			fmt.Print("Done\n")
			wg.Done()
		}()
	}
	wg.Wait()
	waitgroup.Done()
}
