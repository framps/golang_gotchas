package main

// Sample code for a common error using pointers in lists using append
//
// Copyright (C) 2018 framp at linux-tips-and-tricks dot de

import "fmt"

type Element struct {
	Number  int
	Updated bool
}

func main() {

	var ePtr *Element

	list := make([]Element, 0)
	listPtr := make([]*Element, 0)

	for i := 0; i < 5; i++ {
		e := Element{Number: i}
		if i == 2 {
			ePtr = &e // save pointer to element
		}
		fmt.Printf("Adding %d\n", i)
		list = append(list, e)        // element copied into list
		listPtr = append(listPtr, &e) // ptr of element copied in list
	}
	ePtr.Updated = true // update element in list

	fmt.Print("--- List ---\n")
	for i, e := range list {
		fmt.Printf("Element %d, Number: %d, Updated: %t\n", i, e.Number, e.Updated)
	}

	fmt.Print("--- ListPtr ---\n")
	for i, e := range listPtr {
		fmt.Printf("Element %d, Number: %d, Updated: %t\n", i, e.Number, e.Updated)
	}

}
