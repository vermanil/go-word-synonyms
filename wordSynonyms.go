package main

import (
	"fmt"
	"os"
	"encoding/xml"
	"net/http"
	"io/ioutil"
)

const (
	synonymsURl = " https://fraze.it/api/syn/%s/en/360b4932-8c19-464e-93bd-cfb9e62a46ab"
)

type synos struct {
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	Type string `xml:"type"`
	Values []Val `xml:"val"`
}

type Val struct {
	Synonyms []Synonym `xml:"synonyms"`
}

type Synonym struct {
	Syn []string `xml:"synonym"`	
}

var demoXml = []byte(`<synos><total>2</total><entry><type>verb</type><val><synonyms><synonym>actuate</synonym><synonym>operate</synonym></synonyms></val></entry><entry><type>noun</type><val><synonyms><synonym>might</synonym><synonym>force</synonym><synonym>strength</synonym><synonym>potency</synonym><synonym>authority</synonym><synonym>energy</synonym><synonym>vigor</synonym><synonym>vigour</synonym><synonym>capacity</synonym><synonym>ability</synonym></synonyms></val></entry></synos>`)

func formatUrl(query string) string {
	return fmt.Sprintf(synonymsURl, query)
}


func main() {

	query := os.Args[1:][0]
	url := formatUrl(query)	
	resp, err := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(bytes)
	resp.Body.Close()

	if err != nil {
	    fmt.Println(err)
	}

	fmt.Println(url)
	var syn synos

	xml.Unmarshal(bytes, &syn)
	range_ := len(syn.Entries)
	for a :=0; a < range_; a++ {
		fmt.Println(syn.Entries[a].Type)
		fmt.Println(syn.Entries[a].Values[0].Synonyms[0].Syn)
	}
}