package main

import "fmt"

type Player struct {
	name  string
	age   int
	city  string
	state string
}

func main() {
	play1 := Player{
		name:  "Sam",
		age:   25,
		city:  "Hisar",
		state: "Haryana",
	}
	// var play2 Player
	// play2 = Player{"Thomas", 29, "Ludhiana", "Punjab"}

	fmt.Println("name", play1.name)
	fmt.Println("age", play1.age)
	fmt.Println("city", play1.city)
	play1.age = 64
	fmt.Println("age", play1.age)

}