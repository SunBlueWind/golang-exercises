package main

import "fmt"

func dedup(strings []string) []string {
	i := 0
	for _, s := range strings {
		if strings[i] != s {
			i++
			strings[i] = s
		}
	}
	return strings[:i+1]
}

func main() {
	s := []string{"one", "two", "two", "three", "three", "three", "one", "four", "five", "five"}
	s = dedup(s)
	fmt.Println(s)
	s = []string{"one", "two", "four", "five", "1"}
	s = dedup(s)
	fmt.Println(s)
	s = []string{"one", "one", "one", "one", "one", "one"}
	s = dedup(s)
	fmt.Println(s)
}
