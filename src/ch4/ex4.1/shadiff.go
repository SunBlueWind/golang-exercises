package main

import (
	"crypto/sha256"
	"fmt"
)

func countBits(b byte) int {
	count := 0
	for b > 0 {
		count++
		b = b & (b - 1)
	}
	return count
}

func shaDiff(sha1, sha2 [sha256.Size]byte) int {
	diff := [sha256.Size]byte{}
	for i := range diff {
		diff[i] = sha1[i] ^ sha2[i]
	}
	count := 0
	for _, b := range diff {
		count += countBits(b)
	}
	return count
}

func main() {
	fmt.Println(shaDiff(sha256.Sum256([]byte("q")), sha256.Sum256([]byte("q"))))
	fmt.Println(shaDiff(sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X"))))
	fmt.Println(shaDiff(sha256.Sum256([]byte("q")), sha256.Sum256([]byte("p"))))
	fmt.Println(shaDiff(sha256.Sum256([]byte("asdflsdwe")), sha256.Sum256([]byte("p22341"))))
}
