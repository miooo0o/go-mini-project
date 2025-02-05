package main

import "fmt"

func main() {

	fmt.Println("type 1")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	fmt.Println("type 2")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("type 3")
	for i := range 3 {
		fmt.Println("range", i)
	}

	fmt.Println("type 4")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("type 5")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
