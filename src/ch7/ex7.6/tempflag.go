package main

import (
	"flag"
	"fmt"

	"ch7/ex7.6/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
