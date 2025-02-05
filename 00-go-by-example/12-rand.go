package main

import (
	"fmt"
	"math/rand"
)

func randomNumber() int {
	var x int = rand.Intn(400)
	return x
}

func main() {

	x := randomNumber()

	fmt.Printf("this is %d, ", x)
	// if x >= 0 && x <= 100 {
	// 	fmt.Println("between 0 and 100.")
	// } else if x >= 101 && x <= 200 {
	// 	fmt.Println("between 101 and 200.")
	// } else if x >= 201 && x <= 250 {
	// 	fmt.Println("between 201 and 250.")
	// } else {
	// 	fmt.Println("more than 250.")
	// }

	switch {
	case x <= 100:
		fmt.Println("between 0 and 100.")
	case x >= 101 && x <= 200:
		fmt.Println("between 101 and 200.")
	case x >= 201 && x <= 250:
		fmt.Println("between 201 and 250.")
	default:
		fmt.Println("more than 250.")
	}
}
