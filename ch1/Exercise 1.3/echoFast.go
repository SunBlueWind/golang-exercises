package main

import (
	"fmt"
	"os"
	"strings"
)

func echoFast() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
