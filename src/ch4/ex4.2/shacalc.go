package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "need at least 2 cli args")
		return
	}
	s := os.Args[1]
	if len(os.Args) == 3 {
		if f := os.Args[2]; f == "384" {
			fmt.Printf("%x\n", sha512.Sum384([]byte(s)))
		} else if f == "512" {
			fmt.Printf("%x\n", sha512.Sum512([]byte(s)))
		} else {
			fmt.Fprintf(os.Stderr, "unrecognized flag %s\n", f)
		}
	} else {
		fmt.Printf("%x\n", sha256.Sum256([]byte(s)))
	}
}
