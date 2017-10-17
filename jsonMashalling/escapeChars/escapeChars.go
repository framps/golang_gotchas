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

	type simpleInterfaceType interface{}
	type simpleStruct struct {
		Value string
	}

	var simpleStructVariable simpleStruct = simpleStruct{Value: "Fo\"o\tTAB"}
	var simpleInterfaceVariable simpleInterfaceType

	var simpleStructBlob = []byte(`{ "Value": "Fo\"o\tTAB", "Null" : null }`)

	err := json.Unmarshal(simpleStructBlob, &simpleInterfaceVariable)
	panicIfError(err)
	m := simpleInterfaceVariable.(map[string]interface{})
	v := m["Value"].(string)
	n := m["Null"]
	fmt.Printf("*** Blob unmarshaled into interface ***\n%s\n---\n%#v\n---\nValue: '%s' - Null: %#v\n---\n", simpleStructBlob, simpleInterfaceVariable, v, n)

	err = json.Unmarshal(simpleStructBlob, &simpleStructVariable)
	panicIfError(err)
	fmt.Printf("*** Blob unmarshaled into struct ***\n%s\n---\n%#v\n---\nValue: '%s'\n---\n", simpleStructBlob, simpleStructVariable, simpleStructVariable.Value)

	blob, err := json.Marshal(simpleInterfaceVariable)
	panicIfError(err)
	fmt.Printf("*** Marshaled from interface ***\n%#v\n---\n%s\n---\n", simpleInterfaceVariable, blob)

	simpleStructVariable.Value = "Some\"Value\t followed by NL\nand next line"

	blob, err = json.Marshal(simpleStructVariable)
	panicIfError(err)
	fmt.Printf("*** Marshaled from struct ***\n%#v\n---\n%s\n---\n", simpleStructVariable, blob)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(simpleStructVariable.Value)
	fmt.Printf("*** Encoded struct value ***\n%#s\n---\n'%s'\n", simpleStructVariable.Value, b)

	b = new(bytes.Buffer)
	json.NewEncoder(b).Encode(simpleStructVariable)
	fmt.Printf("*** Encoded struct ***\n%#s\n---\n'%s'\n", simpleStructVariable.Value, b)

}
