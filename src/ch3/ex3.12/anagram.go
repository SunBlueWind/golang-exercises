package main

import (
	"sort"
	"strings"
)

func isAnagram(s1, s2 string) bool {
	a1 := strings.Split(s1, "")
	a2 := strings.Split(s2, "")
	sort.Strings(a1)
	sort.Strings(a2)
	return strings.Join(a1, "") == strings.Join(a2, "")
}

func main() {
	println(isAnagram("abc", "cab"))
	println(isAnagram("abcdeedcba", "aabbccddee"))
	println(isAnagram("aabbccdde", "eebbccdda"))
	println(isAnagram("", "eebbccdda"))
}
