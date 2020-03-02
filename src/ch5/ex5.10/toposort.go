package main

import "fmt"

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
}

func toposort(m map[string][]string) []string {
	var order []string
	var seen = make(map[string]bool)
	var visitAll func([]string)
	visitAll = func(items []string) {
		for _, i := range items {
			if seen[i] {
				continue
			}
			seen[i] = true
			visitAll(m[i])
			order = append(order, i)
		}
	}
	for k := range m {
		visitAll([]string{k})
	}
	return order
}

func main() {
	for i, course := range toposort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
