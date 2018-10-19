// Package main prints its command line arguments to stdout
package main

import (
	"flag"
	"os"
)

func main() {
	slow := flag.Bool("slow", false, "run the slower version")
	out := os.Stdout
	inputs := os.Args[1:]
	flag.Parse()
	if *slow {
		echoSlow(out, inputs)
	} else {
		echoFast(out, inputs)
	}
}
