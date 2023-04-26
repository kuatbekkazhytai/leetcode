package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 1}
	k := 1
	fmt.Println(countKDifference(nums, k))
}

func countKDifference(nums []int, k int) int {
	count := 0

	for i := range nums {
		if len(nums[i:]) == 1 {
			break
		}
		count += getDiffCount(nums[i:], k)
	}

	return count
}

func getDiffCount(nums []int, k int) int {
	count := 0
	first := nums[0]

	for i := 1; i < len(nums); i++ {
		if first-nums[i] == k || first-nums[i] == -1*k {
			count++
		}
	}

	return count
}
