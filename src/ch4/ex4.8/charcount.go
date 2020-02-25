package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

const (
	letter = iota
	digit
	symbol
	lower
	upper
	punc
	space
)

func main() {
	count := [space + 1]int{}
	in := bufio.NewReader(os.Stdin)
	invalid := 0

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			count[letter]++
		}
		if unicode.IsDigit(r) {
			count[digit]++
		}
		if unicode.IsSymbol(r) {
			count[symbol]++
		}
		if unicode.IsLower(r) {
			count[lower]++
		}
		if unicode.IsUpper(r) {
			count[upper]++
		}
		if unicode.IsPunct(r) {
			count[punc]++
		}
		if unicode.IsSpace(r) {
			count[space]++
		}
	}
	fmt.Printf("type\tcount\n")
	fmt.Printf("Letter\t%d\n", count[letter])
	fmt.Printf("Digit\t%d\n", count[digit])
	fmt.Printf("Symbol\t%d\n", count[symbol])
	fmt.Printf("Lower\t%d\n", count[lower])
	fmt.Printf("Upper\t%d\n", count[upper])
	fmt.Printf("Punct\t%d\n", count[punc])
	fmt.Printf("Space\t%d\n", count[space])

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
