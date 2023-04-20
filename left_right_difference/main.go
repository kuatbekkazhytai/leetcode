package main

import "fmt"

func main() {
	nums := []int{10, 4, 8, 3}
	fmt.Println(leftRightDifference(nums))
}

func leftRightDifference(nums []int) []int {
	res := make([]int, 0)

	for i := range nums {
		left := nums[:i]
		right := nums[i+1:]

		if sum(left) > sum(right) {
			res = append(res, sum(left)-sum(right))
		} else {
			res = append(res, sum(right)-sum(left))
		}
	}

	return res
}

func sum(nums []int) int {
	res := 0
	for _, j := range nums {
		res = res + j
	}
	return res
}
