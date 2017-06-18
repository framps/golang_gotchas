package main

// Globalize message by using github.com/nicksnyder/go-i18n/i18n
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/i18n"
)

func main() {

	i18n.MustLoadTranslationFile("en-us.all.json")
	i18n.MustLoadTranslationFile("de-de.all.json")
	i18n.MustLoadTranslationFile("fr-fr.all.json")
	i18n.MustLoadTranslationFile("zh.all.json")

	const defaultLocale = "en-US"

	var locales = []struct {
		loc string
	}{
		{"en-US"},
		{"de-DE"},
		{"fr-FR"},
		{"ar-AR"},
		{"zh"},
	}

	for _, l := range locales {

		fmt.Printf("*** Locale %s\n", l.loc)
		T, _ := i18n.Tfunc(l.loc, defaultLocale)
		fmt.Printf(T("settings_title") + "\n")
		fmt.Printf("%s\n", T("hello_world", map[string]interface{}{"Person": "framp"}))
	}
}
