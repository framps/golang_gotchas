package main

// Sample how mashalling of errors works
//
// See github.com/framps/golang_gotchas for latest code
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

import (
	"encoding/json"
	"fmt"
	"os"
)

// See for details: https://github.com/golang/go/issues/10748

func main() {

	funk := "Open: "
	_, err := os.Open("dummy")
	fmt.Printf("%s Explicit error: ==> %v\n", funk, err)
	m, _ := json.Marshal(err)
	fmt.Printf("%s marshall: ==> %v\n", funk, string(m))
	ms, _ := json.Marshal(err.Error())
	fmt.Printf("%s marshall string: ==> %v\n", funk, string(ms))
	/*
	   Open:  Explicit error: ==> open dummy: no such file or directory
	   Open:  marshall: ==> {"Op":"open","Path":"dummy","Err":2}
	   Open:  marshall string: ==> "open dummy: no such file or directory"
	*/
	funk = "Explicit error: "
	err = fmt.Errorf("%s", "Some error")
	fmt.Printf("Explicit error: ==> %v\n", err)
	m, _ = json.Marshal(err)
	fmt.Printf("%s marshall: ==> %v\n", funk, string(m))
	ms, _ = json.Marshal(err.Error())
	fmt.Printf("%s marshall string: ==> %v\n", funk, string(ms))
	/*
	  Explicit error: ==> Some error
	  Explicit error:  marshall: ==> {}
	  Explicit error:  marshall string: ==> "Some error"
	*/
}
