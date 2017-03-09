package main

// Print the different marshall/unmarshal results of different structures into/from JSON
//
// 1) struct
// 2) struct array
// 3) struct list
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// --- different structs to marshall

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

// --- initialize structs to unmarshall

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

func prettyPrint(body []byte) *bytes.Buffer {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, body, "", "   ")
	if err != nil {
		panic(err)
	}
	return &prettyJSON
}

var structSamples = []struct {
	Name        string
	StructValue interface{}
}{
	{"SimpleStruct", simpleStruct},
	{"SimpleStructArray", simpleStructArray},
	{"SimpleStructList", simpleStructList},
}

var simpleStructBlob = []byte(`
		{"Name": "Foo", "Value": "Bar"}
	`)

var simpleStructArrayBlob = []byte(`[
	{"Name": "Foo1", "Value": "Bar2"},
	{"Name": "Foo2", "Value": "Bar2"}
		]
	`)

var blobSamples = []struct {
	Name      string
	BlobValue interface{}
}{
	{"SimpleStructBlob", simpleStruct},
	{"SimpleStructArrayBlob", simpleStructArrayBlob},
}

func main() {

	fmt.Println("--------------- Go struct to JSON (Marshall) ----------------")

	for _, s := range structSamples {

		buffer, _ := json.Marshal(s.StructValue)
		fmt.Printf("%s: %s\n", s.Name, prettyPrint(buffer))

	}

	fmt.Println("--------------- JSON to GO struct (Unmarshall) ----------------")

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

}
