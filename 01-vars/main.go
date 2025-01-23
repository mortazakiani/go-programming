package main

import "fmt"

func main() {
	//set  first  value for variable
	var  age int
	var  name string
	fmt.Println(age)
	fmt.Println(name)

	//set value for variable
	age = 10
	println(age)

	//set values when define arrayes
	var ages = 10
	var names = "morteza"
	fmt.Println(ages)
	fmt.Println(names)

	//define vars in one line
	var agess,sum int 
	println(agess,sum)

	//set values
	agess=10
	sum=1000
	println(agess,sum)


	//set   variable and values as a groups
	var(
		agesss = 10
		namess ="kiani"
		sums= 9999
	)
	println(agesss,namess,sums)
	
	//faster way
	count := 100
	fmt.Print(count)

	//you can defin several vars together
	count1,count2,count3 := 10,20,30
	fmt.Println(count1,count2,count3)
	//you should have a atleast  one new  var  on left side
	var count4= 5 
	count4,count5,count6 := 10,20,30
	fmt.Println(count4,count5,count6)
}