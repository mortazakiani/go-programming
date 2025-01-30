package main

import "fmt"

type Calculator interface {
	getSpeedsum() int
}

type Car struct {
	name  string
	speed int
}

func (c Car) getSpeedsum() int {
	return c.speed
}
func Printsln(c Calculator)  {
	fmt.Println(c.getSpeedsum())
}

func check(i interface{})  {//can check the type of input variables
	v,ok:= i.(int)
	fmt.Println(v,ok)
}

func main() {
	car1 := Car{
		name:  "BMW",
		speed: 200,
	}

	Printsln(car1)

	a:="morteza"
	b:= 123456789
	check(a)
	check(b)
	
}