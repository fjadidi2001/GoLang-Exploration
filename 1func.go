package main

import (
	"fmt"
)

func summation(num1, num2 float64) float64 {
	var sum = num1 + num2
	return sum

}
func main() {
	num1, num2 := 56, 76.82
	sum := summation(float64(num1), num2)
	fmt.Println("sum is :", sum)

}
