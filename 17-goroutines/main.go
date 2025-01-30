package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello")
}

func numbers()  {
	for i := 0; i<= 5; i++ {
		time.Sleep(250 *time.Millisecond)//we want to see how concurency manged 
		fmt.Print(i)
		}
}
func myName()  {
	myname := "morteza"
	for i := 0; i< len(myname); i++ {
		time.Sleep(450 *time.Millisecond)//we want to see how concurency manged 
		fmt.Print(string(myname[i]))
		}
}

func main() {
	go hello()//we says a thread handel  this  part and go  for aonother staemnet
	go myName()//we says a thread handel  this  part and go  for aonother staemnet
	go numbers()//we says a thread handel  this  part and go  for aonother staemnet
	time.Sleep(3500 * time.Millisecond)//wait for threejobs to  end  if not last commend  execute and programm will finish
	fmt.Println("main")//print this and end  the  program
}