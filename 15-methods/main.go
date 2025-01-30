package main

import (
	"fmt"
)

type User struct {
	name string
	age int
	Cars
}
type Cars struct {
	name string
}


func (u User)print() {//same name  but diffrent  input
	fmt.Println(u.name)
}
func (c Cars)printsln() {
	fmt.Println(c)
}


func (u User)changeName(newName string) {//change without  refrence
	u.name =newName
}
func (u *User)changeAge(newAge int) {//change with refrence
	u.age =newAge
}



func main() {
	user :=User{
		name: "morteza",
		age: 25,
		Cars: Cars{
			name: "mustang",
		},
	}
	user.print()

	//you can see if we use without refrence the struct values  does not  change 	
	fmt.Println(user)
	user.changeName("mehdi")
	fmt.Println(user)
	user.changeAge(29)
	fmt.Println(user)

	user.printsln()//i can directly call  the car strucs  type  methods  


}