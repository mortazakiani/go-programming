package main

import "fmt"

//use  pointer.jpg to  understand what's happend
func main() {
	a := 123
	var b *int = &a //its b pointer that  store address of a var  and  type of b  depend type of a
	fmt.Println(b)


	c := 123
	var d *int
	if d == nil {//empty pointer  value is nill
		fmt.Println(d)
		d = &c
		fmt.Println(d)
	}
	

	
	e := new(int)//when create pointers store address  of value 0 becuse its  int
	fmt.Println(*e)
	fmt.Println(e)




	f :=123
	g := &f
	fmt.Println(f)//you can access the value of var if access  thre address of that 
	fmt.Println(g)

	change(g)
	fmt.Println(f)



	change2(f)
	fmt.Println(g)
}
func change(value *int)  {//can chnage  a value by geting  address of variable 
		*value =456
}

func change2(value int) *int {// get  return a address of varibale
	return &value
}
