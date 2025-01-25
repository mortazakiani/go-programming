package main

import "fmt"

func main() {
	//for  strcture
	for i := 0; i < 5; i++ {
		if i > 3 {
			break//break  from loop
		}
		fmt.Println(i)
	}
	for i := 0; i < 5; i++ {
		if i == 4 {
			continue//except from this 
		}
		fmt.Println(i)
	}

	//infinit  loops
	for i := 1; i  > 0; i++ {
		fmt.Println("infinit")
	}
}