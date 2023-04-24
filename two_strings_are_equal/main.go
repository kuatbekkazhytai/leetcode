package main

import "fmt"

func main() {
	w1 := []string{"ab", "c"}
	w2 := []string{"a", "bc"}
	fmt.Println(arrayStringsAreEqual(w1, w2))
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var s1, s2 string
	for _, j := range word1 {
		s1 += j
	}
	for _, j := range word2 {
		s2 += j
	}

	return s1 == s2
}
