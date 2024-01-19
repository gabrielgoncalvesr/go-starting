package main

import "fmt"

func main() {

	// hello world

	fmt.Println("Hello, Ninjas")

	// strings

	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"

	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory

	var numOne int8 = 21
	var numTwo int8 = -128
	var numThree uint = 25

	fmt.Println(numOne, numTwo, numThree)

	// float

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 23232321312312321321313123.7
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
