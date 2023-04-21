package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(differenceOfSum(nums))
}

func differenceOfSum(nums []int) int {
	var elementSum, digitSum int
	for _, j := range nums {
		elementSum += j
		digitSum += getDigitSum(j)
	}

	if elementSum > digitSum {
		return elementSum - digitSum
	}

	return digitSum - elementSum
}

func getDigitSum(n int) int {
	if n < 10 {
		return n
	}
	s := 0
	for n > 0 {
		s += n % 10
		n /= 10
	}
	return s
}
