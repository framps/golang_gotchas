// Use regex to replace a double quote string containing \n with another contents
//
// Two solutions are provided:
// V1: Buggy version which uses regex.ReplaceAllString
// V2: Working version which uses regex.FindAllString and strings.Replace
//
// run buggy version
// go test regex_test.go -run /V=1
// run working version
// go test regex_test.go -run /V=2
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// replace a double quote string with \n (first version - does't work :-()

func convertMultiplineToHereDoc(value string, regexString string) string {
	re := regexp.MustCompile(regexString)
	result := re.ReplaceAllString(value, "<$1>")
	return result
}

// replace a double quote string with \n (second version - works :-) )

func convertMultiplineToHereDoc2(value string, regexString string) string {
	reg := regexp.MustCompile(regexString)
	matchedStrings := reg.FindAllString(value, -1) // find all double quote strings
	result := value
	for _, m := range matchedStrings {
		if strings.ContainsAny(m, "\n\r") { // if there is a \n included
			n := fmt.Sprintf("<%s>", m[1:len(m)-1])   // new string
			result = strings.Replace(result, m, n, 1) // replace old string with new string
		}
	}
	return result
}

var tests = []struct {
	input    string
	expected string
}{

	{"a \"name\" = \"value\"",
		"a \"name\" = \"value\""},
	{"\"a \\t \\n \n b c \n d e\"",
		"<a \\t \\n \n b c \n d e>"},
	{"a \"name\" = \"val\\\"ue\"",
		"a \"name\" = \"val\\\"ue\""},
	{"a \"name\" = \"val\\tue\"",
		"a \"name\" = \"val\\tue\""},
	{"a \"name\" = \"val\ntue\"",
		"a \"name\" = <val\ntue>"},
	{"a \"name\" = \"val\n\tue\nwert\"",
		"a \"name\" = <val\n\tue\nwert>"},
	{"a \"name\" = \"val\n\tu\\\"e\nwert\"",
		"a \"name\" = <val\n\tu\\\"e\nwert>"},
	{"a \"name\" = \n \"val\n\tu\\\"e\nwert\"",
		"a \"name\" = \n <val\n\tu\\\"e\nwert>"},
	{"a \"name\" = \"val\n\tu\\\"e\nwert\" , \n \"val\n\tu\\\"e\nwert\"",
		"a \"name\" = <val\n\tu\\\"e\nwert> , \n <val\n\tu\\\"e\nwert>"},
	{"a \"name\" = \"val\n\tu\\\"e\nwert\" , \n [ \"val\n\tu\\\"e\nwert\" ] , \"name\" = \"val\n\tu\\\"e\nwert\"",
		"a \"name\" = <val\n\tu\\\"e\nwert> , \n [ <val\n\tu\\\"e\nwert> ] , \"name\" = <val\n\tu\\\"e\nwert>"},
}

func TestFoo(t *testing.T) {
	t.Run("V=1", func(t *testing.T) { ReplaceV1(t) })
	t.Run("V=1", func(t *testing.T) { ReplaceV2(t) })
}

// if input is "a" \n b "c\nd"
// code incorrectly returns < \n b > but should return <c\nd>

func ReplaceV1(t *testing.T) {

	// match string with doublequotes which includes a newline"
	regexString := "(?m)\"((?:[^\"\\\\]|\\\\.)*\n(?:[^\"\\\\]|\\\\.)*)\""

	for _, tt := range tests {
		result := convertMultiplineToHereDoc(tt.input, regexString)
		// t.Logf("\nRegex: %s\nInput:\n%s\nOutput:\n%s\n", regexString, tt.input, result)
		assert.Equal(t, tt.expected, result)
	}
}

// if input is "a" \n b "c\nd"
// code correctly returns <c\nd>

func ReplaceV2(t *testing.T) {

	// match any string with doublequotes
	regexString := "(?m)\"((?:[^\"\\\\]|\\\\.)*)\""

	for _, tt := range tests {
		result := convertMultiplineToHereDoc2(tt.input, regexString)
		// t.Logf("\nRegex: %s\nInput:\n%s\nOutput:\n%s\n", regexString, tt.input, result)
		assert.Equal(t, tt.expected, result)
	}
}
