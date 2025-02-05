package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", " as ")
	switch i {
	case 1:
		fmt.Println("once")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Printf("I am a bool %T\n", t)
		case int:
			fmt.Printf("I am an int %T\n", t)
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whoAmI(true)
	whoAmI(1)
	whoAmI("hey")
	whoAmI(i)
}
