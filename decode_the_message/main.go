package main

import "fmt"

func main() {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"

	fmt.Println(decodeMessage(key, message))

}

func decodeMessage(key string, message string) string {
	parsedKey := parseKey(key)
	newStr := ""
	for _, j := range []rune(message) {
		if j == ' ' {
			newStr += " "
			continue
		}
		for k, l := range []rune(parsedKey) {
			if j == l {
				newStr += string('a' + k)
			}

		}
	}

	return newStr
}

func parseKey(key string) string {
	newStr := ""
	runes := []rune(key)
	m := make(map[rune]rune)
	for _, j := range runes {
		if j >= 97 && j <= 122 {
			if _, ok := m[j]; !ok {
				m[j] = j
				newStr += string(j)
			}
		}
	}

	return newStr
}
