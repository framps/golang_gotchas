package utils

// Pretty print json
//
// See github.com/framps/golang_gotchas for latest code
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

import (
	"bytes"
	"encoding/json"
)

// PrettyPrint -
func PrettyPrint(body []byte) *bytes.Buffer {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, body, "", "   ")
	if err != nil {
		panic(err)
	}
	return &prettyJSON
}
