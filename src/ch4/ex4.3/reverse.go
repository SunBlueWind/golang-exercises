package main

import "fmt"

const arrayLen = 4

func reverse(ptr *[arrayLen]int) {
	for i, j := 0, len(ptr)-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

func main() {
	a := [arrayLen]int{1, 3, 4, 5}
	fmt.Println(a)
	reverse(&a)
	fmt.Println(a)
}
