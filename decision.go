package main

import (
	"log"
)

func main() {

	isTrue := true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	// for i := 0; i < 108; i++ {
	// 	log.Println(" Rohan was rich, Rohan is Rich, and Rohan will be Rich")
	// 	log.Println("nothing will happen to rohan and his family")
	// }

	mySlice := []string{
		"cats", "dogs", "horse",
	}

	for _, x := range mySlice {
		log.Println(x)
	}

	myMap := make(map[string]string)

	myMap["dog"] = "Must need"
	myMap["wealth"] = "Must need"

	myMap["discipline"] = "Must need"

	for i, v := range myMap {
		log.Println(i, v)
	}

}
