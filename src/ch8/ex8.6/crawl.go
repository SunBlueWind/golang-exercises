package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

var depth = flag.Int("depth", 3, "depth of links fetched")

type link struct {
	url   string
	depth int
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func makeLinks(urls []string, depth int) []link {
	var links []link
	for _, url := range urls {
		links = append(links, link{url, depth})
	}
	return links
}

func main() {
	flag.Parse()
	worklist := make(chan []link)
	var n int

	n++
	go func() { worklist <- makeLinks(flag.Args(), 0) }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				if link.depth > *depth {
					continue
				}
				n++
				go func(url string, depth int) {
					worklist <- makeLinks(crawl(url), depth+1)
				}(link.url, link.depth)
			}
		}
	}
}
