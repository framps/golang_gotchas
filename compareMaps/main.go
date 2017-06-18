package main

// Comparison of two maps and whether they are equal

// Keep in mind -> https://golang.org/doc/effective_go.html#maps

// Quote:
// "An attempt to fetch a map value with a key that is not present
// in the map will return the zero value for the type of the entries in the map.
// For instance, if the map contains integers, looking up a non-existent key will return 0."

// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

import "fmt"

func equal(x, y map[string]int, wrongTest bool) bool {

	if len(x) != len(y) {
		return false
	}

	if wrongTest {
		for k, xv := range x {
			if xv != y[k] { // most obvious test but read the quote from effective go !
				return false
			}
		}
		return true
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func main() {

	mA := map[string]int{
		"A": 0,
	}
	mB := map[string]int{
		"B": 42,
	}

	fmt.Printf("map A: %+v\nmap B: %+v\n", mA, mB)

	fmt.Printf("Retrieving non existent element mB[A]: %d\n", mB["A"])

	fmt.Printf("Cmp result (wrong): %t\n", equal(mA, mB, true))
	fmt.Printf("Cmp result (OK): %t\n", equal(mA, mB, false))

	// map A: map[A:0]
	// map B: map[B:42]
	// Retrieving non existent element mB[A]: 0
	// Cmp result (wrong): true
	// Cmp result (OK): false

}
