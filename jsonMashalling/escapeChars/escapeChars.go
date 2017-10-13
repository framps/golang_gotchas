package main

// Evaluate escape char handling in GO
//
// See github.com/framps/golang_gotchas for latest code
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	var simpleInterfaceStruct interface{}
	var simpleStruct struct {
		Value string
	}

	var simpleStructBlob = []byte(`{ "Value": "Fo\"o\tTAB", "Null" : null }`)

	err := json.Unmarshal(simpleStructBlob, &simpleInterfaceStruct)
	panicIfError(err)
	m := simpleInterfaceStruct.(map[string]interface{})
	v := m["Value"].(string)
	n := m["Null"]
	fmt.Printf("*** Unmarshaled into interface ***\n%s\n---\n%#v\n---\nValue: '%s' - Null: %#v\n---\n", simpleStructBlob, simpleInterfaceStruct, v, n)

	err = json.Unmarshal(simpleStructBlob, &simpleStruct)
	panicIfError(err)
	fmt.Printf("*** Unmarshaled into struct ***\n%s\n---\n%#v\n---\nValue: '%s'\n---\n", simpleStructBlob, simpleStruct, simpleStruct.Value)

	blob, err := json.Marshal(simpleInterfaceStruct)
	panicIfError(err)
	fmt.Printf("*** Marshaled from interface ***\n%#v\n---\n%s\n---\n", simpleInterfaceStruct, blob)

	simpleStruct.Value = "Some\"Value\t followed by NL\nand next line"

	blob, err = json.Marshal(simpleStruct)
	panicIfError(err)
	fmt.Printf("*** Marshaled from struct ***\n%#v\n---\n%s\n---\n", simpleStruct, blob)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(simpleStruct.Value)
	fmt.Printf("*** Encoded ***\n%#s\n---\nWith DoubleQuotes: '%s'\n---\nWithout DoubleQuotes: '%s'\n", simpleStruct.Value, b, b.String()[1:b.Len()-2])
}
