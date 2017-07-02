package main

// Print the different marshal/unmarshal results of different structures into/from JSON
//
// 1) GO struct -> JSON and JSON string -> struct
// 2) GO struct array -> JSON and JSON string -> struct array
// 3) GO struct list -> JSON and JSON string -> struct list
//
// See github.com/framps/golang_gotchas for latest code
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

import (
	"encoding/json"
	"fmt"

	"github.com/framps/golang_gotchas/utils"
)

// ------------------------------------
// --- different structs to marshal ---
// ------------------------------------

// SimpleStruct -
type SimpleStruct struct {
	Name  string
	Value string
}

// SimpleStructArray -
type SimpleStructArray []SimpleStruct

// SimpleStructList -
type SimpleStructList map[string]SimpleStruct

var simpleStruct = SimpleStruct{
	Name:  "Foo",
	Value: "Bar",
}

// -----------------------------------
// --- initialize structs to unmarshal
// -----------------------------------

var simpleStructList = SimpleStructList{
	"Element1": SimpleStruct{
		Name:  "Foo1",
		Value: "Bar1",
	},
	"Element2": SimpleStruct{
		Name:  "Foo2",
		Value: "Bar2",
	},
}

var simpleStructArray = SimpleStructArray{
	SimpleStruct{
		Name:  "Foo1",
		Value: "Bar1",
	},
	SimpleStruct{
		Name:  "Foo2",
		Value: "Bar2",
	},
}

var structSamples = []struct {
	Name        string
	StructValue interface{}
}{
	{"SimpleStruct", simpleStruct},
	{"SimpleStructArray", simpleStructArray},
	{"SimpleStructList", simpleStructList},
}

// --------------------------------------
// --- different strings to unmarshal ---
// --------------------------------------

var simpleStructBlob = []byte(`
		{ "Name": "Foo", "Value": "Bar" }
	`)

var simpleStructArrayBlob = []byte(`[
	{ "Name": "Foo1", "Value": "Bar2" },
	{ "Name": "Foo2", "Value": "Bar2" }
		]
	`)

var simpleStructListBlob = []byte(` {
	"Element1" : {
				"Name": "Foo1", "Value": "Bar2"
			},
	"Element2" : {
				"Name": "Foo2", "Value": "Bar2"
			}
	}
	`)

var blobSamples = []struct {
	Name      string
	BlobValue interface{}
}{
	{"SimpleStructBlob", simpleStruct},
	{"SimpleStructArrayBlob", simpleStructArrayBlob},
	{"SimpleStructListBlob", simpleStructListBlob},
}

func main() {

	fmt.Println("--------------- Go struct to JSON (Marshal) ----------------")
	fmt.Println()

	for _, s := range structSamples {

		buffer, _ := json.Marshal(s.StructValue)
		fmt.Printf("%s: %s\n", s.Name, utils.PrettyPrint(buffer))

	}

	fmt.Println("--------------- JSON to GO struct (Unmarshal) ----------------")
	fmt.Println()

	simpleStruct := new(SimpleStruct)
	err := json.Unmarshal(simpleStructBlob, simpleStruct)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %+v\n", "SimpleStructBlob", *simpleStruct)

	simpleStructArray := new(SimpleStructArray)
	err = json.Unmarshal(simpleStructArrayBlob, simpleStructArray)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %+v\n", "SimpleStructArrayBlob", *simpleStructArray)

	simpleStructList := new(SimpleStructList)
	err = json.Unmarshal(simpleStructListBlob, simpleStructList)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %+v\n", "SimpleStructListBlob", *simpleStructList)

}

// Execution result:

/*
framp@obelix:~/go/src/github.com/framps/golang_gotchas/jsonMashalling (master)$ go run jsonMarshalling.go
--------------- Go struct to JSON (Marshal) ----------------

SimpleStruct: {
   "Name": "Foo",
   "Value": "Bar"
}
SimpleStructArray: [
   {
      "Name": "Foo1",
      "Value": "Bar1"
   },
   {
      "Name": "Foo2",
      "Value": "Bar2"
   }
]
SimpleStructList: {
   "Element1": {
      "Name": "Foo1",
      "Value": "Bar1"
   },
   "Element2": {
      "Name": "Foo2",
      "Value": "Bar2"
   }
}
--------------- JSON to GO struct (Unmarshal) ----------------

SimpleStructBlob: {Name:Foo Value:Bar}
SimpleStructArrayBlob: [{Name:Foo1 Value:Bar2} {Name:Foo2 Value:Bar2}]
SimpleStructListBlob: map[Element1:{Name:Foo1 Value:Bar2} Element2:{Name:Foo2 Value:Bar2}]
*/
