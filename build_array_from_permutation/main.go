package main

import "fmt"

func main() {
	nums := []int{5, 0, 1, 2, 3, 4}
	fmt.Println(buildArray(nums))
}

func buildArray(nums []int) []int {
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[nums[i]])
	}

	return ans
}
