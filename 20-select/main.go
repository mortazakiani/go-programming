package main

import (
	"fmt"
	"time"
)

//its  look like  switch  command
func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "server2"
}


func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)
	//its cheeck  which var  has  a value and show that  
	select {
	case s1 := <- output1:
		fmt.Print(s1)
	case s2 := <- output2:
		fmt.Print(s2)
	default:
		fmt.Println("anyone")
	}
	

}