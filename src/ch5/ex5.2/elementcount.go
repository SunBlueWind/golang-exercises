package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	count := make(map[string]int)
	visit(doc, count)
	for elem, c := range count {
		fmt.Printf("<%s>:\t%d\n", elem, c)
	}
}

func visit(n *html.Node, count map[string]int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	visit(n.FirstChild, count)
	visit(n.NextSibling, count)
}
