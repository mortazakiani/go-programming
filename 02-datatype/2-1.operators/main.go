package main

import "fmt"

func main() {
	c, x, y := 0, 1, 1
	c = x + y
	fmt.Println(c)
	c = x - y
	fmt.Println(c)
	c = x * y
	fmt.Println(c)
	c = x / y
	fmt.Println(c)
	c = x % y
	fmt.Println(c)
	c += 10 // c is c = c + 10
	fmt.Println(c)
	//and it's true for others - ,* , / , %
}