package main

import (
	"fmt"
	"os"
)

func echoSlow() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += arg + sep
		sep = " "
	}
	fmt.Println(s)
}
