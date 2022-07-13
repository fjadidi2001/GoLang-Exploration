//package main
//import "fmt"
//func A&P(Base,Height float64)(float64,float64){
//	var Arr=we
//
//}
//func main(){
//	var Base float64=15
//	var Height float64=23.14
//	fmt.Printf("")
//}

package main

import (
	"fmt"
)

func Calculation(Height, Base float64) (float64, float64) {

	var area = (Base * Height) / 2
	var perimeter = Base * 3
	return perimeter, area

}
func main() {
	area, perimeter := Calculation(10.8, 5.6) // perimeter is discarded

	fmt.Printf("Area type %T,Area vale %d,Area Binary %b\n", area, perimeter)
	println(area, perimeter)
}
