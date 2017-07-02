package main

// Copyright (C) 2017 framp at linux-tips-and-tricks dot de
//
// Sample how to debug http client by using io/ioutil
//
// See github.com/framps/golang_tutorial for latest code
//
// This code is based and enhanced on sample code from http://speakmy.name/2014/07/29/http-request-debugging-in-go/

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}

func main() {

	var body []byte
	var response *http.Response
	var request *http.Request

	url := "http://maps.googleapis.com/maps/api/geocode/json?address=Stuttgart%2C+K%C3%B6nigstrasse%2C+1"

	request, err := http.NewRequest("GET", url, nil)
	if err == nil {
		request.Header.Add("Content-Type", "application/json")
		fmt.Printf("********** Request\n")
		debug(httputil.DumpRequestOut(request, true))
		response, err = (&http.Client{}).Do(request)
	}

	if err == nil {
		defer response.Body.Close()
		fmt.Printf("********** Request\n")
		debug(httputil.DumpResponse(response, true))
		body, err = ioutil.ReadAll(response.Body)
	}

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	_ = body
}
