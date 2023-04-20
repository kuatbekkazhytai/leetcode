package main

import "fmt"

func main() {
	str := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}

	fmt.Println(mostWordsFound(str))
}

func mostWordsFound(sentences []string) int {
	max := 0
	for _, j := range sentences {
		if countSpaces(j) > max {
			max = countSpaces(j)
		}
	}

	return max
}

func countSpaces(str string) int {
	space := 1

	for _, j := range str {
		if j == ' ' {
			space++
		}
	}
	return space
}
