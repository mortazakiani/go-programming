package main

import "fmt"

func main() {
	a := map[string]string{//you  can choose you type and build a map 
		"name":     "morteza",
		"lastname": "kiani",
	}
	fmt.Println(a)
	a["location"]="tad"//add an element
	fmt.Println(a)



	b:= map[string]string{
		"name":     "morteza",
		"lastname": "kiani",
	}

	if value,ok :=b["name"];ok {//you can  use thsi  if condition  for one time not anywhere 
		fmt.Println(value)
	}


	c:= map[string]string{// how to delete  element 
		"name":     "morteza",
		"lastname": "kiani",
	}
	fmt.Println(c)
	delete(c,"name")
	fmt.Println(c)

	d:= map[string]string{// how to  get  element  of the map 
		"name":     "morteza",
		"lastname": "kiani",
	}
	fmt.Println(d)
	for i, v := range d {
	fmt.Println(i,v)	
	}

	e:= map[string]string{// maps are  refrence types if you change value it changed in refrence
		"name":     "morteza",
		"lastname": "kiani",
	}
	test := e
	test["name"] ="mehdi"
	fmt.Println(e)

}