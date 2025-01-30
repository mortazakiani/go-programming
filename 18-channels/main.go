package main

import "fmt"

// in channels  when  the programs is done that  all fuction that we call we go routine is done
//and we cannect go  rountines together
func hello(data chan bool) {
	fmt.Println("hi")
	data <- true
}


func senddata(data chan<- int)  {
	data<-123//wreit only channel we can not read 
}
func getdata(data <-chan int)  {
	
	a := <-data//readolny only channel we can not read 
	fmt.Println(a)
}


func closechannel(data chan int)  {
	for i := 0; i < 10; i++ {//close channel 
		data <-i	
	}
	close(data)
}


func cal1(numbers int,ch1 chan int){
 numbers *= 2
 ch1 <- numbers
}
func cal2(numbers int,ch2 chan int){
	numbers += 2
	ch2 <- numbers
}
func main() {

data := make(chan bool)
go hello(data)
a:= <-data
fmt.Println(a)
fmt.Println("main")

numbers :=10
c := make(chan int)
d := make(chan int)
go cal1(numbers,c)
go cal2(numbers,d)
s1,s2 := <-c,<-d
fmt.Println(s1+s2)

}