package main

import "fmt"

func main() {
	fmt.Println(sumOfMultiples(7))
}

func sumOfMultiples(n int) int {
	count := 0

	for i := 3; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			count += i
		}
	}

	return count
}
