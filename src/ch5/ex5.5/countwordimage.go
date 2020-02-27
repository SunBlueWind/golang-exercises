package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countwordimage: %v\n", err)
			continue
		}
		fmt.Printf("[%s]\n", url)
		fmt.Printf("words:\t%d\n", words)
		fmt.Printf("images:\t%d\n", images)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return 0, 0
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	} else if n.Type == html.TextNode {
		words += countWords(n.Data)
	}
	childWords, childImages := countWordsAndImages(n.FirstChild)
	siblingWords, siblingImages := countWordsAndImages(n.NextSibling)
	words += (childWords + siblingWords)
	images += (childImages + siblingImages)
	return
}

func countWords(s string) (words int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	return
}
