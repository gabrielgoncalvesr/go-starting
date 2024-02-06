package main

import "fmt"

func class5() {

	x := 0

	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	fmt.Println("-")

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	fmt.Println("-")

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	for i := 0; i < len(names); i++ {
		fmt.Println("value of i is:", names[i])
	}

	fmt.Println("-")

	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	fmt.Println("-")

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
	}

	fmt.Println("-")

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new string"
	}

	fmt.Println(names)
}
