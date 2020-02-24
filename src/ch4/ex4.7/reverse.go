package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(bytes []byte) {
	for i, j := 0, len(bytes); i < j; {
		firstR, firstSize := utf8.DecodeRune(bytes[i:j])
		lastR, lastSize := utf8.DecodeLastRune(bytes[i:j])
		copy(bytes[i+lastSize:j-firstSize], bytes[i+firstSize:j-lastSize])
		utf8.EncodeRune(bytes[i:], lastR)
		utf8.EncodeRune(bytes[j-firstSize:], firstR)
		i += lastSize
		j -= firstSize
	}
}

func main() {
	b := []byte("hello wolrd 你好啊 this is great 哈哈")
	fmt.Printf("%s\n", b)
	reverse(b)
	fmt.Printf("%s\n", b)
}
