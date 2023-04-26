package main

import "fmt"

func main() {
	s := "B0R0G0R9R0B0G0"
	fmt.Println(countPoints(s))
}

func countPoints(rings string) int {
	count := 0
	m := make(map[string][]string)
	var key, value string
	for i := 1; i < len(rings); i = i + 2 {
		key = string(rings[i])
		value = string(rings[i-1])

		if !inSlice(m[key], value) {
			m[key] = append(m[key], value)
		}
	}

	for _, j := range m {
		if len(j) == 3 {
			count++
		}
	}

	return count
}

func inSlice(str []string, value string) bool {
	for _, j := range str {
		if value == j {
			return true
		}
	}

	return false
}
