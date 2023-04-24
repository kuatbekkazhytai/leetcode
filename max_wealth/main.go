package main

import "fmt"

func main() {
	nums := [][]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println(maximumWealth(nums))
}

func maximumWealth(accounts [][]int) int {
	max := 0
	for i := 0; i < len(accounts); i++ {
		if count(accounts[i]) > max {
			max = count(accounts[i])
		}
	}

	return max
}

func count(nums []int) int {
	sum := 0

	for _, j := range nums {
		sum += j
	}

	return sum
}
