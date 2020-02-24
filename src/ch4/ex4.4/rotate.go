package main

import "fmt"

func rotate(a []int, i int) {
	i %= len(a)
	tmp := make([]int, i)
	copy(tmp, a[:i])
	copy(a[:len(a)-i], a[i:])
	copy(a[len(a)-i:], tmp)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(a, 6)
	fmt.Println(a)
}
