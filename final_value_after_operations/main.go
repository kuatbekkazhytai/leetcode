package main

import "fmt"

func main() {
	str := []string{"X++", "++X", "--X", "X--"}

	fmt.Println(finalValueAfterOperations(str))
}

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, j := range operations {
		if j == "--X" || j == "X--" {
			x--
		} else {
			x++
		}
	}
	return x
}
