package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s) % 3
	buf.WriteString(s[0:n])
	for ; n+3 <= len(s); n += 3 {
		if n > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[n : n+3])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("12345"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("12345678909876543212345678"))
	fmt.Println(comma("1234567890987654321234567890"))
}
