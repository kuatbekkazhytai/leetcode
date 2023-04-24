package main

import "fmt"

func main() {
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
	key := "color"
	value := "silver"
	fmt.Println(countMatches(items, key, value))
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var count int

	for _, j := range items {
		if ruleKey == "type" && j[0] == ruleValue {
			count++
		}
		if ruleKey == "color" && j[1] == ruleValue {
			count++
		}
		if ruleKey == "name" && j[2] == ruleValue {
			count++
		}
	}

	return count
}
