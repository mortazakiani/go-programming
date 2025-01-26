package main

import "fmt"

func main() {
	a := [3]int{80, 90, 100}// create  arraye  
	var b []int = a[1:2]//set slice  of  arraye 
	fmt.Println(b)


	c := []int{80,70,60,50,40}//create  arraye without specfic numberof elemnt
	fmt.Println(c)

//it cab show in slices we just take a piece of arraye and we changed it 
	d := [6]int{80,70,60,50,40}
	var f []int = d[0:6]
	fmt.Println("before ",d)
	for i := range f {
		f[i]++
	}
	fmt.Println("after ",d)
	fmt.Println("len ",len(f)," capacity",cap(f))
}