package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(replaceDigits("a1b2c3d4e"))
}

func replaceDigits(s string) string {
	var res string
	for i := 0; i < len(s)-1; i = i + 2 {
		res += string(s[i])
		n, _ := strconv.Atoi(string(s[i+1]))
		res += string(n + int(s[i]))
		fmt.Println(res)
	}

	if len(s)%2 == 1 {
		res += string(s[len(s)-1])
	}

	return res
}
