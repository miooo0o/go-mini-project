package main

import (
	"fmt"
)

func main() {
	const name = "Kim"
	const age = 22

	fmt.Printf("%s is %d years old\n", name, age)
	fmt.Printf("%s is %v years old. \t and the type is %T and %T", name, age, name, age)

}
