// Package main prints its command line arguments to stdout
package main

import (
	"flag"
)

func main() {
	slow := flag.Bool("slow", false, "run the slower version")
	flag.Parse()
	if *slow {
		echoSlow()
	} else {
		echoFast()
	}
}
