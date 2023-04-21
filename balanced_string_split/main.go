package main

import "fmt"

func main() {
	s := "LLLLRRRR"
	fmt.Println(balancedStringSplit(s))
}

func balancedStringSplit(s string) int {
	var countR, countL, counter int

	for _, j := range s {
		if j == 'L' {
			countL++
		} else {
			countR++
		}
		if countL == countR && countL > 0 {
			counter++
			countL = 0
			countR = 0
		}
	}

	return counter
}
