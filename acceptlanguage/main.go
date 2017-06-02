package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/text/language"
)

// matcher is a language.Matcher configured for all supported languages.
var matcher = language.NewMatcher([]language.Tag{
	language.Make("en-GB"),
	language.BritishEnglish,
	language.Norwegian,
	language.German,
})

// handler is a http.HandlerFunc.
func handler(w http.ResponseWriter, r *http.Request) {
	t, q, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	// We ignore the error: the default language will be selected for t == nil.
	tag, _, _ := matcher.Match(t...)
	fmt.Printf("%5v (t: %6v; q: %3v; err: %v)\n", tag, t, q, err)
}

func main() {
	for _, al := range []string{
		"nn;q=0.3, en-us;q=0.8, en,",
		"gsw, en;q=0.7, en-US;q=0.8",
		"gsw, nl, da",
		"invalid",
	} {
		// Create dummy request with Accept-Language set and pass it to handler.
		r, _ := http.NewRequest("GET", "example.com", strings.NewReader("Hello"))
		r.Header.Set("Accept-Language", al)
		handler(nil, r)
	}

}
