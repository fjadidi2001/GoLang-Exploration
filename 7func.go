package main

import "fmt"

func calculation(side1, side2, base1, base2, Height float64) (float64, float64, float64) {
	var area = base1 * Height
	var perimeter1 = side1 + side2 + base1 + base2
	var perimeter2 = (2 * base1) + (2 * side1)
	return area, perimeter1, perimeter2
}

func main() {
	area, _, _ := calculation(14.23, 14.2333333334, 15, 15.223333, 23.2425)
	fmt.Println(area)
}

//In line no. 13 we use only the area and the _ identifier is used to discard the perimeter.
