package main

import "fmt"

func main() {
	s := "is2 sentence4 This1 a3"
	fmt.Println(sortSentence(s))
}

func sortSentence(s string) string {
	m := make(map[int]string)
	var temp, res string
	runes := []rune(s)
	for _, j := range runes {
		if j == ' ' {
			continue
		}
		if j >= '1' && j <= '9' {
			m[int(j-48)] = temp
			temp = ""
			continue
		}
		temp += string(j)
	}

	for i := 1; i <= len(m); i++ {
		res += m[i] + " "
	}

	return res[:len(res)-1]
}
