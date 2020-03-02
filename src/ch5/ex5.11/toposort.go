package main

import (
	"fmt"
	"log"
)

var prereqs = map[string][]string{
	"algorithm": {"data structures"},
	"calculus":  {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":        {"calculus"},
}

const (
	white = iota
	gray
	black
)

func toposort(m map[string][]string) ([]string, bool) {
	var order []string
	var colors = make(map[string]int)
	for k := range m {
		colors[k] = white
	}
	var visitAll func([]string) bool
	visitAll = func(items []string) bool {
		for _, i := range items {
			if colors[i] == gray {
				// cycle found
				return false
			}
			if colors[i] != white {
				continue
			}
			colors[i] = gray
			if !visitAll(m[i]) {
				return false
			}
			order = append(order, i)
			colors[i] = black
		}
		return true
	}
	for k := range m {
		if !visitAll([]string{k}) {
			return nil, false
		}
	}
	return order, true
}

func main() {
	chain, ok := toposort(prereqs)
	if !ok {
		log.Fatal("toposort: cycle found")
	}
	for i, course := range chain {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
