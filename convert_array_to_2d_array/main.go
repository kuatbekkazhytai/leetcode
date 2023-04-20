package main

import "fmt"

func main() {
	nums := []int{8, 8, 8, 8, 2, 4, 4, 2, 4}

	fmt.Println(findMatrix(nums))
}

func findMatrix(nums []int) [][]int {
	res := make([][]int, len(nums))

	for _, j := range nums {

		exists, index := inSlice(res, j)
		if exists {
			res[index+1] = append(res[index+1], j)
			continue
		}

		res[0] = append(res[0], j)
	}
	for true {
		if len(res[len(res)-1]) == 0 {
			res = res[:len(res)-1]
			continue
		}
		break
	}

	return res

}

func inSlice(nums [][]int, needle int) (bool, int) {
	for i := len(nums) - 1; i >= 0; i-- {
		for _, l := range nums[i] {
			if needle == l {
				return true, i
			}
		}
	}

	return false, 0
}
