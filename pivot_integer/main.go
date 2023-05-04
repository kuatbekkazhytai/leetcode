package main

import "fmt"

func main() {
	fmt.Println(pivotInteger(8))
}

func pivotInteger(n int) int {
	for i := 1; i <= n; i++ {
		if isPivot(i, n) {
			return i
		}
	}

	return -1
}

func isPivot(i, n int) bool {
	var l, r int

	for j := 1; j <= i; j++ {
		l += j
	}

	for j := i; j <= n; j++ {
		r += j
	}

	return l == r
}
