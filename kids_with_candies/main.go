package main

import "fmt"

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	fmt.Println(kidsWithCandies(candies, extraCandies))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var res []bool
	for _, j := range candies {
		res = append(res, isGreatest(j+extraCandies, candies))
	}

	return res
}

func isGreatest(comp int, candies []int) bool {
	for _, j := range candies {
		if comp < j {
			return false
		}
	}

	return true
}
