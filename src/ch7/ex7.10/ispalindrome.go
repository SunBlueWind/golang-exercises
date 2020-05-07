package main

import (
	"fmt"
	"sort"
)

// IsPalindrome reports whether s is a palindrome
func IsPalindrome(s sort.Interface) bool {
	i, j := 0, s.Len()-1
	for i < j {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(sort.StringSlice([]string{"a", "b", "c"})))
	fmt.Println(IsPalindrome(sort.StringSlice([]string{"a", "b", "a"})))
	fmt.Println(IsPalindrome(sort.StringSlice([]string{"a", "b", "b", "a"})))
}
