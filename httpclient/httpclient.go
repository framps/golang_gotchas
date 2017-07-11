package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {

	fmt.Println("Starting SERVER1 ...")
	svr1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("SERVER1: Received request. Going to sleep for one hour")
		time.Sleep(time.Hour)
		fmt.Println("SERVER1: Processing request")
	}))
	defer svr1.Close()

	fmt.Println("CLIENT1: making request")
	/* --- Will sleep for one hour :-(
	http.Get(svr1.URL)
	*/

	// Use a timeout of 3 seconds instead
	var netClient = &http.Client{
		Timeout: time.Second * 3,
	}

	response, err := netClient.Get(svr1.URL)
	if err != nil {
		fmt.Printf("CLIENT1: finished request with error. %s\n", err.Error())
	} else {
		fmt.Printf("CLIENT1: finished request. Response statuscode: %d\n", response.StatusCode)
	}

	fmt.Println("Starting SERVER2 ...")
	svr2 := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("SERVER2: Hello Client")
	}))
	svr2.Config.ReadTimeout = 3 * time.Second
	svr2.Config.WriteTimeout = 10 * time.Second
	svr2.Start()
	defer svr2.Close()

	fmt.Println("CLIENT2: making request")
	http.Get(svr2.URL)

}
