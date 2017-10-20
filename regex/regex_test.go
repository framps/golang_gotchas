package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func convertMultiplineToHereDoc(value string) string {
	re := regexp.MustCompile(`\"([^"\n]+?\n+[^"\n]+?)\"`)
	result := re.ReplaceAllString(value, "<\n$1>\n")
	return result
}

// table drivern test - usual test pattern in go
var tests = []struct {
	input    string
	expected string
}{

	// empty strings
	{"\"111\"\"\"222 \n, 333\"\"444",
		"\"111\"\"\"222 \n, 333\"\"444"},
	// empty strings with nl
	{"\"111\"\"\n\"222 \n, 333\"\n\"444",
		"\"111\"<\n\n>\n222 \n, 333<\n\n>\n444"},
	/*
		{"\"var\" = \"v11\n" +
			"v12\"222 ,\n" +
			"111\"v21\n" +
			"v22\"222",
			"\"var\" = <\n" +
				"v11\n" +
				"v12\n" +
				">\n" +
				"222 ,\n" +
				"111<\n" +
				"v21\n" +
				"v22\n" +
				">\n" +
				"222"},
	*/
}

func TestReplace(t *testing.T) {
	for _, tt := range tests {
		result := convertMultiplineToHereDoc(tt.input)
		t.Logf("\n>>>:\n%s\n<<<:\n%s\n", tt.input, result)
		assert.Equal(t, tt.expected, result)
	}
}
