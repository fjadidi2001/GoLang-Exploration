package main

import "fmt"

func calculation(base1, base2, Hight, side1, side2 float64) (float64, float64) {
	var area = ((base1 + base2) / 2) * Hight
	var perimeter = base1 + base2 + side1 + side2
	return area, perimeter
}
func main() {
	area, perimeter := calculation(16, 14, 17, 10.23, 10.23)
	fmt.Println(area, perimeter)
	fmt.Printf("area %T,perimeter %T\n", area, perimeter)

}
