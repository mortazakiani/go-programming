package main

import (
	"fmt"
	"sync"
	"time"
)


func prtintsln(i int ,wg *sync.WaitGroup)  {
	fmt.Println("start",i)
	time.Sleep(2 * time.Second)
	fmt.Println("end",i)
	wg.Done()
}



func main() {
	ch := make(chan int, 3)//we can said it jsut  accept 3 elemnt
	ch <- 10
	ch <- 100
	ch <- 1000
	//if we read a value its gone
	fmt.Println(<-ch)
	 fmt.Println(<-ch)
	 fmt.Println(<-ch)
	//we can see  when we read a  value it deceasre  to 2
	// fmt.Println("len",len(ch))
	// fmt.Println("cap",cap(ch))
	// fmt.Println("read",<-ch)
	// fmt.Println("len",len(ch))
	// fmt.Println("cap",cap(ch))

	number :=3
	var  wg sync.WaitGroup

	for i := 0; i < number; i++ {
		wg.Add(1)
		go prtintsln(i,&wg)
	}
	
	wg.Wait()
	fmt.Println("finish")
}