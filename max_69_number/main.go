package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximum69Number(9669))
}

func maximum69Number(num int) int {
	runes := []rune(strconv.Itoa(num))
	for i, j := range runes {
		if j == '6' {
			runes[i] = '9'
			break
		}
	}

	r, _ := strconv.Atoi(string(runes))

	return r
}
