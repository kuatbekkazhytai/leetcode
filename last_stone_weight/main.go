package main

import "fmt"

func main() {
	stones := []int{2, 2}
	fmt.Println(lastStoneWeight(stones))
}

func lastStoneWeight(stones []int) int {
	l := len(stones)

	for i := 0; i < l; i++ {
		if len(stones) == 1 {
			break
		}
		x, y := twoMaxNums(stones)

		if x[1] == y[1] {
			ind1 := x[0]
			ind2 := y[0]
			if ind1+1 < len(stones) {
				stones = append(stones[:ind1], stones[ind1+1:]...)
			} else {
				stones = stones[:len(stones)-1]
			}
			if len(stones) == 1 {
				stones = stones[:0]
				break
			}
			if ind2+1 < len(stones) {
				stones = append(stones[:ind2], stones[ind2+1:]...)
			} else {
				stones = stones[:len(stones)-1]
			}

		} else if x[1] < y[1] {
			stones[y[0]] = y[1] - x[1]
			stones = append(stones[:x[0]], stones[x[0]+1:]...)
		}
	}

	if len(stones) == 0 {
		return 0
	}

	return stones[0]
}

func twoMaxNums(stones []int) ([]int, []int) {
	max1 := []int{0, 0}
	max2 := []int{0, 0}
	for i, j := range stones {
		if j > max1[1] {
			max2 = max1
			max1 = []int{i, j}
			continue
		}
		if j > max2[1] {
			max2 = []int{i, j}
		}
	}

	return max2, max1
}
