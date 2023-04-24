package main

import "fmt"

func main() {
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
	key := "color"
	value := "silver"
	fmt.Println(countMatches(items, key, value))
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var isType, isColor, isName bool
	var count int
	if ruleKey == "type" {
		isType = true
	} else if ruleKey == "color" {
		isColor = true
	} else {
		isName = true
	}
	for _, j := range items {
		if isType && j[0] == ruleValue {
			count++
		}
		if isColor && j[1] == ruleValue {
			count++
		}
		if isName && j[2] == ruleValue {
			count++
		}
	}

	return count
}

