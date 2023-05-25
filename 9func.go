package main

import (
	"fmt"
)

func multi(num1, num2 float64) float64 {
	var multi = num1 * num2
	return multi

}
func main() {
	num1, num2 := 56, 76.82
	multi := multii(float64(num1), num2)
	fmt.Println("multi is :", multi)

}
