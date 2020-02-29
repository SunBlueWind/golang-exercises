package elementbyid

import (
	"golang.org/x/net/html"
)

var depth int

func forEachNode(n *html.Node, pre, post func(*html.Node) bool) *html.Node {
	if pre != nil {
		stop := pre(n)
		if stop {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		n := forEachNode(c, pre, post)
		if n != nil {
			return n
		}
	}
	if post != nil {
		post(n)
	}
	return nil
}

func checkIDFunc(id string) func(*html.Node) bool {
	return func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					return true
				}
			}
		}
		return false
	}
}

// ElementByID returns the first HTML element with the specified
// id attribute
func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, checkIDFunc(id), nil)
}
