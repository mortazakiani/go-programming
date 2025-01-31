package main

import (
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		//return 0, errors.New("calculate faild because radius less than zero")//simple return string  for error
		return 0, fmt.Errorf("calculate faild because radius %f less than zero",radius)//we can pass  a variabel inside a string of error

	}
	return math.Pi*radius*radius,nil
}
func main() {
	
	radius :=-20.00
	area,err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("araea is : ",area)
}