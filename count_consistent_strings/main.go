package main

import "fmt"

func main() {
	allowed := "abc"
	str := []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
	fmt.Println(countConsistentStrings(allowed, str))
}

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	newAllowed := toSlice(allowed)
	for _, j := range words {
		if isConsistent(j, newAllowed) {
			count++
		}
	}

	return count
}

func toSlice(allowed string) []string {
	var sl []string
	for _, j := range allowed {
		sl = append(sl, string(j))
	}

	return sl
}

func isConsistent(check string, allowed []string) bool {
	for _, j := range check {
		inSliceBool := inSlice(string(j), allowed)
		if !inSliceBool {
			return false
		}
	}

	return true
}

func inSlice(s string, str []string) bool {
	matches := 0
	for _, j := range str {
		if j == s {
			matches++
			break
		}
	}

	return matches > 0
}
