package main

import (
	"fmt"
	"os"
	"strings"
)

func stringsJoin(sep string, strs ...string) string {
	b := strings.Builder{}
	for i, s := range strs {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(s)
	}
	return b.String()
}

func main() {
	if len(os.Args) > 1 {
		s := stringsJoin(os.Args[1], os.Args[2:]...)
		fmt.Printf("resulting string is: %s\n", s)
	}
}
