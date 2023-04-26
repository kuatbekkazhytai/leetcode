package main

import "fmt"

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	res := ""
	temp := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			res = temp + " " + res
			temp = ""
			continue
		}
		temp += string(s[i])
	}
	res = temp + " " + res

	return res[:len(res)-1]
}
