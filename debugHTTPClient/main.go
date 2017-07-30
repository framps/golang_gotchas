package main

// Copyright (C) 2017 framp at linux-tips-and-tricks dot de
//
// Sample how to debug http client by using io/ioutil and httputil
// Use param -debug to enable http client debugging
//
// See github.com/framps/golang_gotchas for latest code
//
// This code is based and was enhanced
// from sample code http://speakmy.name/2014/07/29/http-request-debugging-in-go/

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

var debugEnabled bool

func debug(description string, data []byte, err error) {
	if debugEnabled {
		if err == nil {
			fmt.Printf("********** %s: %s\n\n", description, data)
		} else {
			log.Fatalf("********** %s: %s\n\n", description, err)
		}
	}
}

func main() {

	var body []byte
	var response *http.Response
	var request *http.Request

	flag.BoolVar(&debugEnabled, "debug", false, "Enable HTTP client debug")
	flag.Parse()

	url := "http://maps.googleapis.com/maps/api/geocode/json?address=Stuttgart%2C+K%C3%B6nigstrasse%2C+1"

	request, err := http.NewRequest("GET", url, nil)
	if err == nil {
		request.Header.Add("Content-Type", "application/json")
		b, e := httputil.DumpRequestOut(request, true)
		debug("Request", b, e)
		response, err = (&http.Client{}).Do(request)
	}

	if err == nil {
		defer response.Body.Close()
		b, e := httputil.DumpResponse(response, true)
		debug("Response", b, e)
		body, err = ioutil.ReadAll(response.Body)
	}

	if !debugEnabled {
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		} else {
			r := strings.NewReplacer(" ", "", "\n", "", "\t", "") // strip whitespaces
			log.Printf("%s", r.Replace(string(body)))
		}
	}
}
