package main

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/i18n"
)

func main() {

	i18n.MustLoadTranslationFile("en-us.all.json")
	i18n.MustLoadTranslationFile("de-de.all.json")
	i18n.MustLoadTranslationFile("fr-fr.all.json")

	T, _ := i18n.Tfunc("de-DE")
	fmt.Printf("I18N text: %s\n", T("settings_title"))
	fmt.Printf("I18N hello: %s\n", T("hello_world"))

	T, _ = i18n.Tfunc("fr-FR")
	fmt.Printf("I18N text: %s\n", T("settings_title"))
	fmt.Printf("I18N hello: %s\n", T("hello_world"))

	T, _ = i18n.Tfunc("en-US")
	fmt.Printf("I18N text: %s\n", T("settings_title"))
	fmt.Printf("I18N hello: %s\n", T("hello_world"))

}
