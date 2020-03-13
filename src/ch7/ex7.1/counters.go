package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type WordCounter int
type LineCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w++
	}
	return int(*w), nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*l++
	}
	return int(*l), nil
}

func main() {
	var w WordCounter
	var l LineCounter
	if len(os.Args) > 1 {
		w.Write([]byte(os.Args[1]))
		fmt.Printf("words: %d\n", w)
		l.Write([]byte(os.Args[1]))
		fmt.Printf("lines: %d\n", l)
	}
}
