package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	// fmt.Println(a, b)

	f := "apple"
	fmt.Println(f)
	f = "1"
	// f = 1
	fmt.Println(f)
}
