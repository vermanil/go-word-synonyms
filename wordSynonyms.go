package main

import (
	"fmt"
	"os"
	// "net/http"
)

func googletranslateUrl(accept_lang string, from_lang string, to_lang string, query string) string {
	return fmt.Sprintf("https://translate.google.com/translate?hl=%s&sl=%s&tl=%s&q=%s", accept_lang, from_lang, to_lang, query)
}

func main() {

	query := os.Args[1:][0]
	accept_lang := "en"
	from_lang := "en"
	to_lang := "hi"
	url := googletranslateUrl(accept_lang, from_lang, to_lang, query)	
	fmt.Println(url)

}