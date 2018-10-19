package main

import (
	"fmt"
	"io"
)

func echoSlow(w io.Writer, inputs []string) {
	s, sep := "", ""
	for _, arg := range inputs {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(w, s)
}
