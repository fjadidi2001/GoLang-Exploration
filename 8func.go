package main

import (
	"fmt"
)

func divi(num1, num2 float64) float64 {
	var div = num1 / num2
	return div

}
func main() {
	num1, num2 := 56, 76.82
	div := divi(float64(num1), num2)
	fmt.Println("div is :", div)

}
