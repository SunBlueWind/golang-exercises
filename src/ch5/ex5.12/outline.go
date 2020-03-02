package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func outline(n *html.Node) {
	var depth int
	startElement := func(n *html.Node) {
		if n.Type == html.CommentNode {
			fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
		}
		if n.Type == html.TextNode {
			trimedText := strings.TrimSpace(n.Data)
			if len(trimedText) > 0 {
				for _, s := range strings.Split(trimedText, "\n") {
					fmt.Printf("%*s%s\n", depth*2, "", s)
				}
			}
		} else if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			for _, attr := range n.Attr {
				fmt.Printf(" %s='%s'", attr.Key, attr.Val)
			}
			if n.FirstChild == nil {
				fmt.Printf("/>\n")
			} else {
				fmt.Printf(">\n")
				depth++
			}
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode && n.FirstChild != nil {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	forEachNode(n, startElement, endElement)
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
		outline(doc)
	}
}
