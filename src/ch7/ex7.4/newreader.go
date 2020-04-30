package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

// Reader is used in NewReader, and it behaves like strings.Reader
type Reader struct {
	s   string
	idx int
}

func (r *Reader) Read(p []byte) (int, error) {
	if r.idx >= len(r.s) {
		return 0, io.EOF
	}
	n := copy(p, r.s[r.idx:])
	r.idx += n
	return n, nil
}

// NewReader behaves like strings.NewReader
func NewReader(s string) *Reader {
	return &Reader{s: s}
}

func main() {
	r := NewReader("<h1>Test</h1>")
	doc, err := html.Parse(r)
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
