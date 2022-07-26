package main

import "log"

func main() {

	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	log.Println(myMap["dog"])

	var mySlie []string

	mySliceUsingShortHand := []int{1, 3, 6, 9, 1, 0, 8}

	mySlie = append(mySlie, "rohan", "Shukla")

	log.Println(mySlie)

	log.Println(mySliceUsingShortHand[0:3])

}
