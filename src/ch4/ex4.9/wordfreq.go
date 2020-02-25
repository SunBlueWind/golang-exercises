package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		freq[scanner.Text()]++
	}
	fmt.Printf("word\tfreq\n")
	for word, freq := range freq {
		fmt.Printf("%s\t%d\n", word, freq)
	}
}
