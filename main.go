package main

import (
	"log"
	"time"
)

type User struct {
	FullName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func (m *myStruct) printFirstName() string {

	return m.FirstName
}

type myStruct struct {
	FirstName string
}

func main() {

	user := User{
		FullName:    "Trevor",
		PhoneNumber: "1212121",
	}

	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "John op",
	}
	log.Println(user.FullName, user.PhoneNumber, user.BirthDate)

	log.Println(myVar.FirstName, myVar2.FirstName)

	log.Println(myVar.printFirstName())

}
