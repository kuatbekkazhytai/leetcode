package main

import "fmt"

func main() {
	s := "yo|uar|e**|b|e***au|tifu|l"
	fmt.Println(countAsterisks(s))
}

func countAsterisks(s string) int {
	count := 0
	isPaired := true
	for _, j := range s {
		if isPaired && j == '*' {
			count++
			continue
		}
		if j == '|' && isPaired {
			isPaired = false
			continue
		}
		if j == '|' && !isPaired {
			isPaired = true
		}
	}

	return count
}
