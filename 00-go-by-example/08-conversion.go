package main

import "fmt"

func main() {
	a := 42
	b := 42.0 // Implicitly float64 (default in Go)
	// b := float64(42.0) // Unnecessary conversion

	fmt.Printf("%v of type %T\n", a, a)
	fmt.Printf("%v of type %T\n", b, b)

	var c float32 = 43.742
	fmt.Printf("%v of type %T\n", c, c) // Explicitly use float32, floating-point

	b = float64(c)
	fmt.Printf("%v of type %T\n", b, b)

}

/*
uint(iota)               // iota value of type uint
float32(2.718281828)     // 2.718281828 of type float32
complex128(1)            // 1.0 + 0.0i of type complex128
float32(0.49999999)      // 0.5 of type float32
float64(-1e-1000)        // 0.0 of type float64
string('x')              // "x" of type string
string(0x266c)           // "♬" of type string
myString("foo" + "bar")  // "foobar" of type myString
string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
(*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
int(1.2)                 // illegal: 1.2 cannot be represented as an int
string(65.0)             // illegal: 65.0 is not an integer constant
*/
