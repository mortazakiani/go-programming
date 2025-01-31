package main

import (
	"fmt"
	"os"
)

func main() {
	contens,err := os.ReadFile("text.txt")
	if err != nil{
		fmt.Println("file reading error",err)
		return
	}
	fmt.Println(string(contens))
	for _, v := range contens {
		fmt.Printf("%c \n",v)
	}

	content,err := os.Create("test.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	l,err := content.WriteString("morteza")
	if err != nil {
		fmt.Println(err)
		content.Close()
		return
	}
	fmt.Println(l)
	err=content.Close()
	if err != nil{
		fmt.Println(err)
		return
	}
}