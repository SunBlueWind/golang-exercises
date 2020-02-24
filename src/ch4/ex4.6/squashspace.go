package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpace(bytes []byte) []byte {
	i, j := 0, 0
	inSpace := false
	for {
		r, size := utf8.DecodeRune(bytes[j:])
		j += size
		if size == 0 {
			break
		}
		if unicode.IsSpace(r) {
			inSpace = true
			continue
		}
		if inSpace {
			n := utf8.EncodeRune(bytes[i:], ' ')
			i += n
		}
		inSpace = false
		n := utf8.EncodeRune(bytes[i:], r)
		i += n
	}
	return bytes[:i]
}

func main() {
	b := []byte("hello  world   !!     this  is a     tab 		你 好	啊	??!!")
	b = squashSpace(b)
	fmt.Printf("%s\n", b)
}
