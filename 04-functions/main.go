package main

import "fmt"

func main() {
	price ,count := 100,2
	//we cal it with the name of the function and parameters that need
	fmt.Println(sum(price,count))
	//we cal it with the name of the function and parameters that need and resturn how many  return we defind
	c1,c2 := sum2(price,count)
	fmt.Println(sum(c1,c2))
	//if we need one the return numbers just use _ and igonre return value
	c3, _ := sum2(price,count)
	fmt.Println(c3)

}
//if we  had just one return value and two parameters
func sum(price,count int) int {
	sum := price * count
	return sum
}

//if we  have multi return value and  multiple parameters to return
func sum2(price,count int) (int,int) {
	sum := price * count
	sum2 := price + count
	return sum,sum2
}