// Package dup prints the name, counts, and lines that appear more than once in the input.
// It can read from the stdin or from a list of files provided on the command line.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		counts["stdin"] = make(map[string]int)
		countLines(os.Stdin, counts["stdin"])
	} else {
		for _, file := range files {
			counts[file] = make(map[string]int)
			countFileLines(file, counts[file])
		}
	}
	printCounts(os.Stdout, counts)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if len(input.Text()) > 0 {
			// does not count empty lines
			counts[input.Text()]++
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dup: scanning: %v\n", err.Error())
	}
}

func countFileLines(name string, counts map[string]int) {
	f, err := os.Open(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup: opening file %q: %v", name, err)
		return
	}
	defer f.Close()
	countLines(f, counts)
}

func printCounts(w io.Writer, counts map[string]map[string]int) {
	for name, countMap := range counts {
		fmt.Fprintf(w, "%v:\n", name)
		for line, count := range countMap {
			if count > 1 {
				fmt.Fprintf(w, "%d\t%s\n", count, line)
			}
		}
		fmt.Println() // to print an empty line between files
	}
}
