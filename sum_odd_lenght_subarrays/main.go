package main

import "fmt"

func main() {
	nums := []int{1, 4, 2, 5, 3}
	fmt.Println(sumOddLengthSubarrays(nums))
}

func sumOddLengthSubarrays(arr []int) int {
	count := 0
	l := len(arr)
	for i := 0; i < l; i++ {
		count += countInSubArrays(arr[i:])
	}

	return count
}

func countInSubArrays(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			count += countArrayDigits(arr[:i+1])
		}
	}

	return count
}

func countArrayDigits(arr []int) int {
	count := 0
	for _, j := range arr {
		count += j
	}

	return count
}
