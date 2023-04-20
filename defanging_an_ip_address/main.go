package main

import "fmt"

func main() {
	str := "255.100.50.0"
	fmt.Println(defangIPaddr(str))
}

func defangIPaddr(address string) string {
	ans := ""
	for _, j := range address {
		if j == '.' {
			ans = ans + "[.]"
		} else {
			ans = ans + string(j)
		}
	}

	return ans
}
