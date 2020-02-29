package main

import (
	"fmt"
	"os"
	"strings"
)

const target = "$foo"

func expand(s string, f func(string) string) string {
	return strings.Replace(s, target, f("foo"), -1)
}

func simpleReplace(_ string) string {
	return "***"
}

func main() {
	for _, s := range os.Args[1:] {
		fmt.Printf("%s -> %s\n", s, expand(s, simpleReplace))
	}
}
