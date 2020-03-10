package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// ElementsByTagName returns all the elements in doc that match
// one of the names in name
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var res []*html.Node
	names := make(map[string]bool)
	for _, n := range name {
		names[n] = true
	}
	if doc.Type == html.ElementNode {
		if names[doc.Data] {
			res = append(res, doc)
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, ElementsByTagName(c, name...)...)
	}
	return res
}

func main() {
	if len(os.Args) > 1 {
		resp, err := http.Get(os.Args[1])
		if err != nil {
			log.Fatalf("http GET failed: %v\n", err)
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatalf("html parse failed: %v\n", err)
		}
		elements := ElementsByTagName(doc, os.Args[2:]...)
		for _, el := range elements {
			fmt.Println(el.Data, el.Attr)
		}
	}
}
