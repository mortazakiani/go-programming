package main

import (
	"fmt"
	"math"
)

func main() {
	const a = "test"
	fmt.Println(a)
	a ="test2" //it cannot to be happend 
	const(
		b=2
		c="morteza"
		d=1.22
	)
	var e = math.Sqrt(4)
	fmt.Println(e)
	const f = math.Sqrt(4) //cannot to be happend beacuse its filled in run time not compile time
	fmt.Println(f)
}