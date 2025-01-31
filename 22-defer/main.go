package main

import "fmt"

func print() {
	fmt.Println("morteza")
}
func value(a int)  {
	fmt.Println(a)
}

func main() {
	defer print()//delayed to  run  the command
	fmt.Println("hello")

	a:=1
	defer value(a)//send  1 but does not  execute  an after  all  execute  with 1 value
	a=10
	fmt.Println(a)
	//its go  to  end  and then from end to  start  run defer comamnd
}