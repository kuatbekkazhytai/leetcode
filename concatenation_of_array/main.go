package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 1}
	fmt.Println(getConcatenation(nums))
}

func getConcatenation(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < 2*len(nums); i++ {
		if i < len(nums) {
			res = append(res, nums[i])
		} else {
			res = append(res, nums[i-len(nums)])
		}
	}

	return res
}
