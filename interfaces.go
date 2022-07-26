package main

import "log"

type Animal interface {
	Says() string
	TotalLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Snake struct {
	Name        string
	Color       string
	PoisonLevel string
}

func main() {

	dog := Dog{
		Name:  "stupid samson",
		Breed: "Husky",
	}

	PrintInfo(dog)

}

func (d Dog) Says() string {
	return "woof"
}

func (d Dog) TotalLegs() int {
	return 4
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.TotalLegs())
}
