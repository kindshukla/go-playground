package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	// myJson := `
	// [
	// 	{
	// 		"id": 1,
	// 		"email": "lazy@gmail.com",
	// 		"age": 12
	// 	},
	// 	{
	// 		"id": 2,
	// 		"email": "lazy2@gmail.com",
	// 		"age": 18
	// 	},
	// 	{
	// 		"id": 3,
	// 		"email": "lazy3@gmail.com",
	// 		"age": 369
	// 	}
	// ]`

	// var dankusers []Person

	// err := json.Unmarshal([]byte(myJson), &dankusers)

	// if err != nil {

	// 	log.Println(err)
	// }

	// log.Printf("dankusers: %v", dankusers)

	var mySlice []Person

	m1 := Person{
		Id:    108,
		Email: "SHUKLA@GMAIL.COM",
		Age:   27,
	}

	mySlice = append(mySlice, m1)

	m2 := Person{
		Id:    107,
		Email: "SHUKLA1@GMAIL.COM",
		Age:   27,
	}

	mySlice = append(mySlice, m2)

	myJson, err := json.MarshalIndent(mySlice, "", "  ")

	if err != nil {
		log.Println("error Marshallin", err)
	}

	fmt.Println(string(myJson))

}
