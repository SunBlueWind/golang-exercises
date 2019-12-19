package main

import (
	"fmt"

	"ch2/ex2.1/tempconv"
)

func main() {
	fmt.Println(tempconv.CToK(tempconv.FreezingC))
	fmt.Println(tempconv.CToK(tempconv.AbsoluteZeroC))
}
