package main

import "fmt"

//if we  have a lot  of  options

func main() {
	num := 10
	switch num {
	case 1:
		fmt.Println(num)
	case 2:
		fmt.Println(num)
	case 3:
		fmt.Println(num)
	case 4:
		fmt.Println(num)
	case 5:
		fmt.Println(num)
	case 6:
		fmt.Println(num)
	case 7:
		fmt.Println(num)
	case 8:
		fmt.Println(num)
	case 9:
		fmt.Println(num)
	default:
		fmt.Println(num)


	}
//if some of the  choices  have common response 
	num2 := 11
	switch num2 {
	case 1,2,3,5:
		fmt.Println(num)
	default	:
		fmt.Println("number is set")
	}
//we can set  nothing in condition and it has true value by defult => switch true{... and our
//cases should contain a false  or true value and when we have true its return that  value
	num3 := 3
	switch  {
	case num3 > 5:
		fmt.Println(num)
	default	:
		fmt.Println("number is setsssss")
	}
}