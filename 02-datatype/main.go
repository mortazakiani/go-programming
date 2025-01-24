package main

import "fmt"

func main() {
	//bool its a variable that can be true or false
	var  a bool = true
	fmt.Println(a)
	var b bool =false
	fmt.Println(b)
	//if both of them is true return true else return fasle
	c :=a && b
	fmt.Println(c)
	//String is a type get some text and set between to ""
	var name string ="morteza"
	fmt.Println(name)
	var name2 string = "kiani"
	fullname :=name+" "+name2
	fmt.Println(fullname)
	/*numerics

		int is get +,- 
		int 
		int8
		int16
		int32
		int64

		unit is get just posetive number
		unit
		unit8
		unit16
		unit32
		unit64


		float32
		float64
		complex64
		complex128

		byte is a alias of unit8
		rune  is alias  of  int32


		compailer will choose between them base on system architect
		if you want  to optimise the space  you can use thes base on your need 
	*/
	d := -5 //is  int or int 64 
	var e uint  = 6  //is unit  type and we should say that in our code
	var f  float32 = 5.324 //is float and get 
	fmt.Println("d is",d,"e is" , e,"f is" ,f)
	//is mathematical use
	g := complex(5,6)
	fmt.Println(g)



}