package main

import "fmt"

func main() {
	jewels := "aA"
	stones := "aAAbbbb"

	fmt.Println(numJewelsInStones(jewels, stones))
}

func numJewelsInStones(jewels string, stones string) int {
	count := 0
	for _, j := range stones {
		if isJewel(j, jewels) {
			count++
		}
	}

	return count
}

func isJewel(j int32, jewel string) bool {
	for _, k := range jewel {
		if k == j {
			return true
		}
	}
	return false
}
