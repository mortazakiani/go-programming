package main

import (
	"fmt"
)

func main() {
	var a [5]int//create empty arraye with 5 elemnts 
	b := [4]int{5,6,7,8}//create a arraye with 4 elemnt (with value)
	c := [...]int {10,20,30,40,50}//create array without specific  lenght of arraye 
	fmt.Println(b,a,c)
	fmt.Println("lenght a is : ",len(a))//get  lenght of arraye
	for i ,v := range b {// navigate  in arraye to return  value and index
		fmt.Println(i,v)
	}
	for _ ,v := range c {//navigate  in arraye to return just value
		fmt.Println(v)
	}
	e := [3][2]string{
		{"a","b"},
		{"c","d"},
		{"e","f"},

	}
	fmt.Println(e)
}