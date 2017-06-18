package main

// Select the appropriate language for HTTP Accept-Header using golang.org/x/text/language
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/text/language"
)

// matcher is a language.Matcher configured for all supported languages.
var matcher1 = language.NewMatcher([]language.Tag{
	language.BritishEnglish,
	language.Norwegian,
	language.German,
})

var matcher2 = dynamicMatcher([]string{"en-GB", "no", "de"})

func dynamicMatcher(languages []string) language.Matcher {
	t := make([]language.Tag, 0, len(languages))
	for _, l := range languages {
		t = append(t, language.Make(l))
	}
	m := language.NewMatcher(t)
	return m
}

// handler is a http.HandlerFunc.
func handler(w http.ResponseWriter, r *http.Request, matcher language.Matcher) {
	t, q, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	// We ignore the error: the default language will be selected for t == nil.
	tag, _, _ := matcher.Match(t...)
	fmt.Printf("%5v (t: %6v; q: %3v; err: %v)\n", tag, t, q, err)
}

func main() {

	for _, m := range []language.Matcher{matcher1, matcher2} {
		for _, al := range []string{
			"nn;q=0.3, en-us;q=0.8, en,",
			"gsw, en;q=0.7, en-US;q=0.8",
			"gsw, nl, da",
			"invalid",
		} {
			// Create dummy request with Accept-Language set and pass it to handler.

			r, _ := http.NewRequest("GET", "example.com", strings.NewReader("Hello"))
			r.Header.Set("Accept-Language", al)
			handler(nil, r, m)
		}
		fmt.Println()
	}

}
