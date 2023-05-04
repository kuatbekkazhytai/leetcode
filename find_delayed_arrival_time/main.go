package main

import "fmt"

func main() {
	fmt.Println(findDelayedArrivalTime(15, 5))
}

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	time := arrivalTime + delayedTime

	if time == 24 {
		return 0
	} else if time > 24 {
		return time - 24
	}

	return time
}
