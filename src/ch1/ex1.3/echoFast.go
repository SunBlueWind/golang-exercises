package main

import (
	"fmt"
	"io"
	"strings"
)

func echoFast(w io.Writer, inputs []string) {
	fmt.Fprintln(w, strings.Join(inputs, " "))
}
