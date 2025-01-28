package main

import (
	"fmt"
	"test/models"
)

type User struct {// its create  a tyep  and  we can use it like a map
	name      string
	lastename string
	age       int
	address Address
}
type Address struct{
	city string
	zip int
}
type Car struct {
	string
	bool
	int
}

func main() {
	user1 := User{
		name:      "morteza",
		lastename: "kiani",
		age:       29,
	}
	fmt.Println(user1)


	user2 := struct {//anonymouse structs
		name      string
		lastename string
		age       int
	}{
		name:      "morteza",
		lastename: "kiani",
		age:       28,
	}
	fmt.Println(user2)
	fmt.Println(user2.age)//jsut use age  or any other element


	var user3  User//use  empty values of their types
	fmt.Println(user3.name)
	fmt.Println(user3.lastename)
	fmt.Println(user3.age)


	car1 := Car{//its named strucs a element with  types 
		string: "jaguar",
		bool: true,
		int: 32,
	}
	fmt.Println(car1)

	user4 := User{
		name:      "morteza",
		lastename: "kiani",
		age:       29,
		address: Address{
			city: "esfahan",
			zip: 8463146856,
		},
	}
	fmt.Println(user4)
	home := models.Home{
		Floors: 2,
		For_sales: true,
	}
	fmt.Println(home)

	
}