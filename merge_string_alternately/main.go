package main

import "fmt"

func main() {
	w1 := "ab"
	w2 := "pqrs"
	fmt.Println(mergeAlternately(w1, w2))
}

func mergeAlternately(word1 string, word2 string) string {
	var str, longer string
	var l, i int
	if len(word1) > len(word2) {
		l = len(word2)
		longer = word1
	} else {
		l = len(word1)
		longer = word2
	}
	for i < l {
		str += string(word1[i])
		str += string(word2[i])
		i++
	}
	str += longer[i:]

	return str
}
