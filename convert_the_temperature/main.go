package main

import (
	"fmt"
)

func main() {
	temp := 36.50
	fmt.Println(convertTemperature(temp))
}

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32}
}
