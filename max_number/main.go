package main

import "fmt"

func main() {
	n := []int{3, 4, 5, 2}

	fmt.Println(maxProduct(n))
}

func maxProduct(nums []int) int {
	var max, second int
	for _, j := range nums {
		if j > max {
			second = max
			max = j
		} else if j > second {
			second = j
		}
	}

	return (max - 1) * (second - 1)
}
