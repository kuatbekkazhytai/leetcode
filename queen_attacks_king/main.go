package main

import "fmt"

func main() {
	queen := [][]int{{5, 6}, {7, 7}, {2, 1}, {0, 7},
		{1, 6}, {5, 1}, {3, 7}, {0, 3}, {4, 0}, {1, 2}, {6, 3}, {5, 0}, {0, 4},
		{2, 2}, {1, 1}, {6, 4}, {5, 4}, {0, 0}, {2, 6}, {4, 5}, {5, 2}, {1, 4},
		{7, 5}, {2, 3}, {0, 5}, {4, 2}, {1, 0}, {2, 7}, {0, 1}, {4, 6}, {6, 1},
		{0, 6}, {4, 3}, {1, 7}}
	king := []int{3, 4}

	fmt.Println(queensAttackTheKing(queen, king))
}

func addHorizontally(queen []int, king []int, res [][]int) [][]int {
	notAdded := true
	if len(res) == 0 {
		return append(res, queen)
	}
	for i, existing := range res {
		if !onDiagonal(existing, king) {
			if queen[1] < king[1] && existing[1] < king[1] {
				if distance(queen, king) < distance(existing, king) {
					res[i] = queen
				}
				notAdded = false
				break
			} else if queen[1] > king[1] && existing[1] > king[1] {
				if distance(queen, king) < distance(existing, king) {
					res[i] = queen
				}
				notAdded = false
				break
			}
		}
	}
	if notAdded {
		res = append(res, queen)
	}

	return res
}

func addVertically(queen []int, king []int, res [][]int) [][]int {
	notAdded := true
	if len(res) == 0 {
		return append(res, queen)
	}
	for i, existing := range res {
		if !onDiagonal(existing, king) {
			if queen[0] < king[0] && existing[0] < king[0] {
				if distance(queen, king) < distance(existing, king) {
					res[i] = queen
				}
				notAdded = false
				break
			} else if queen[0] > king[0] && existing[0] > king[0] {
				if distance(queen, king) < distance(existing, king) {
					res[i] = queen
				}
				notAdded = false
				break
			}
		}
	}
	if notAdded {
		res = append(res, queen)
	}
	if queen[0] == 5 && queen[1] == 4 {
		fmt.Println("res is", res, queen)
	}

	return res
}

func addDiagonal(queen []int, king []int, res [][]int) [][]int {
	notAdded := true
	if len(res) == 0 {
		return append(res, queen)
	}
	for i, existing := range res {
		if onDiagonal(queen, existing) && onDiagonal(existing, king) {
			if queen[0] < king[0] && existing[0] < king[0] {
				if distance(queen, king) < distance(existing, king) {

					res[i] = queen
				}
				notAdded = false
				break
			} else if queen[0] > king[0] && existing[0] > king[0] {
				if distance(queen, king) < distance(existing, king) {
					res[i] = queen
				}
				notAdded = false
				break
			}
		}
	}
	if notAdded {
		res = append(res, queen)
	}

	return res
}

func distance(x []int, y []int) int {
	return (x[0]-y[0])*(x[0]-y[0]) + (x[1]-y[1])*(x[1]-y[1])
}

func onDiagonal(king []int, queen []int) bool {
	alreadyReversed := false
	if king[0]-queen[0] == king[1]-queen[1] {
		return true
	}
	diff1 := king[0] - queen[1]
	diff2 := king[1] - queen[0]
	if diff1 > 0 && diff2 > 0 {
		return false
	}
	if diff1 < 0 {
		alreadyReversed = true
		diff1 = diff1 * -1
	}
	if diff2 < 0 && !alreadyReversed {
		diff2 = diff2 * -1
	}
	if diff1 == diff2 {
		return true
	}

	return false
}

func queensAttackTheKing(queens [][]int, king []int) [][]int {
	var res [][]int
	for _, queen := range queens {
		fmt.Println("queen is", queen, "res is", res)
		if queen[0] == king[0] && !onDiagonal(king, queen) {
			res = addHorizontally(queen, king, res)
			continue
		}
		if queen[1] == king[1] && !onDiagonal(king, queen) {
			res = addVertically(queen, king, res)
			continue
		}
		if onDiagonal(king, queen) {
			res = addDiagonal(queen, king, res)
		}
	}

	return res
}
