package main
import "fmt"

func Calc(Sum, Min, Mul, Div float64) (float64, float64) {
	var area = (Base * Height) / 2
	var perimeter = side1 + side2 + Base
	return area, perimeter
}
func main() {
	area, perimeter := calculation(2, 10.2, 10.5, 10)
	fmt.Println(area, perimeter)
	fmt.Printf("AREA TYPE:%T,PERIMETER TYPE %T\n", area, perimeter)
}
