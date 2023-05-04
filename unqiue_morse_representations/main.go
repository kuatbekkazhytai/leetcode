package main

import "fmt"

func main() {
	words := []string{"a"}

	fmt.Println(uniqueMorseRepresentations(words))

}

func uniqueMorseRepresentations(words []string) int {
	var res []string
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
		"....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.",
		"--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	for _, j := range words {
		word := ""
		for _, l := range j {
			word = word + morse[l-97]
		}
		if !inSlice(res, word) {
			res = append(res, word)
		}
	}

	return len(res)
}

func inSlice(res []string, word string) bool {
	for _, j := range res {
		if j == word {
			return true
		}
	}

	return false
}
