package main

import "fmt"

func main() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}
	fmt.Println(sortPeople(names, heights))
}

func sortPeople(names []string, heights []int) []string {
	for i := 0; i < len(heights); i++ {
		for k := i + 1; k < len(heights); k++ {
			if heights[i] < heights[k] {
				heights[i], heights[k] = heights[k], heights[i]
				names[i], names[k] = names[k], names[i]
			}
		}
	}

	return names
}
