package main

import "fmt"

func main() {
	fmt.Println(countDigits(54))
}

func countDigits(num int) int {
	div := 1000
	val := num
	count := 0
	for div > 0 {
		div = num % 10
		if div == 0 {
			break
		}
		if val%div == 0 {
			count++
		}
		num = num / 10
	}

	return count
}
