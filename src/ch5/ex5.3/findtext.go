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
	visit(doc)
}

func visit(n *html.Node) {
	if n == nil {
		return
	}
	if n.Data == "script" || n.Data == "style" {
		return
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	visit(n.FirstChild)
	visit(n.NextSibling)
}
