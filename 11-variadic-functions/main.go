package main

import "fmt"

func main() {
	numbers := []int{100,200,100,300,400,10}
	find(10,numbers...)//we can call func with slice like this and every one of elemnt  send to function
	find(10,10,100,200,100,300,400,10)//call varadic-func normally 
}
func find(number int, numbers ...int) {//this is varadic-func
	found := false
	for i, v := range numbers {
		if v == number {
			fmt.Println(number," find in index ",i)
			found = true
		}
	}
	if  !found {
		fmt.Println(number,"not found")
	}
}