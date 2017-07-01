package main

// Use gofuncs to execute get requests against a number of urls and calculate responsetime statistics
//
// Usage
//  -loops int
//    	Request loops per url thread (default 3)
//  -sleep duration
//    	Number of milliseconds to sleep between url requests (default 1s)
//  -threads int
//    	Number of threads per URL (default 3)
//  -urls string
//    	urls to send requests to (default "http://www.google.de")
//  -verbose
//    	verbose logging
//
// Example: go run fetchURLs.go -urls "http://www.cnn.com https://www.thesun.co.uk/" -sleep 1ms -threads 100 -loops 10
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// collect statistics per url
type urlStatistics struct {
	url   string
	count int
	sum   float64
	avg   float64
	max   *float64
	min   *float64
}

var verbose bool

func info(format string, args ...interface{}) {
	if verbose {
		fmt.Printf(format, args...)
	}
}

// kick off GET request threads per url
func fetch(url string, ch chan<- urlStatistics, threads int, sleep time.Duration, loops int) {

	stats := urlStatistics{url: url}

	info("fetching %s\n", url)

	var wg sync.WaitGroup
	wg.Add(threads) // populate wait group

	for i := 0; i < threads; i++ {
		go func(i int) { // i has to be passed, otherwise it will be #threads all the time because i is shared by threads
			info("Starting %d: %s\n", i, url)

			var (
				min, max float64
			)

			for i := 0; i < loops; i++ {
				start := time.Now()

				// kick off a GET request
				resp, err := http.Get(url)
				if err != nil {
					panic(err)
				}
				nbytes, err := io.Copy(ioutil.Discard, resp.Body)
				resp.Body.Close()
				if err != nil {
					panic(err)
				}

				// GET request finished, calculate statistics
				secs := time.Since(start).Seconds()
				info("%d: %s - %.2fs  %7d\n", i, url, secs, nbytes)

				stats.count++
				if stats.min == nil {
					min = secs
					stats.min = &min // don't use &secs because then stats.min and stats.max will be identical all the time
				} else if secs < *stats.min {
					min = secs
				}
				if stats.max == nil {
					max = secs
					stats.max = &max
				} else if secs > *stats.max {
					max = secs
				}
				stats.sum += secs

				// throttle requests
				time.Sleep(sleep)
			}
			info("Done %d: %s\n", i, stats.url)
			wg.Done()
		}(i) // see comment above
	}
	wg.Wait() // wait for all threads to finish

	stats.avg = stats.sum / float64(stats.count)
	ch <- stats // return stats
}

func main() {

	// parse invocation parms

	threads := flag.Int("threads", 3, "Number of threads per URL")
	sleep := flag.Duration("sleep", 1000*time.Millisecond, "Number of milliseconds to sleep between url requests")
	loops := flag.Int("loops", 3, "Request loops per url thread")
	flag.BoolVar(&verbose, "verbose", false, "verbose logging")
	urlList := flag.String("urls", "http://www.google.de", "urls to send requests to")
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Printf("No args required\n")
		os.Exit(1)
	}

	fmt.Printf("Threads: %d, Loops: %d, Sleeps: %s\n", *threads, *loops, *sleep)

	start := time.Now()
	ch := make(chan urlStatistics) // used by fetch tasks per url to report stats when done

	urls := strings.Split(*urlList, " ")

	// kick off fetch threads per url
	for _, url := range urls {
		info("Starting to poll %s\n", url)
		go fetch(url, ch, *threads, *sleep, *loops)
	}

	// wait for tasks to finish and print stats
	var sumRequests int
	for i := 0; i < len(urls); i++ {
		s := <-ch
		fmt.Printf("Stats of %s: N: %d, Min: %.2fs, Avg: %.2fs, Max: %.2fs\n", s.url, s.count, *s.min, s.avg, *s.max)
		sumRequests += s.count
	}

	// report some interesting performance stats
	elapsed := time.Since(start).Seconds()
	avgTimePerRequest := elapsed / float64(sumRequests)
	fmt.Printf("%.2fs elapsed, Avg time per request: %.2fs\n", elapsed, avgTimePerRequest)
}
