package utils

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
