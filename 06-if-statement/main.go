package main

import "fmt"

func main() {
	number := 10

	if number == 2 {// it does not have  any pranteses
		fmt.Println("number 2")
	} else { // it should be on the line of where  bracket of if  closed
		fmt.Println("number :",number)
	}
	age:=30
	if age < 20 {
		fmt.Println("he/she is young")
	}else if age <30 {//esle if structure 
		fmt.Println("he/she is not  old")
	} else{
		fmt.Println("he/she is old")
	}


	if price:=30;  price < 20 {// you can define var  in if  staement  but  its valid in the  if block
		fmt.Println("he/she is young")
	}else if price <30 {
		fmt.Println("he/she is not  old")
	} else{
		fmt.Println("he/she is old")
	}
	// here  we again define  price
	price :=40
	fmt.Println(price)
}