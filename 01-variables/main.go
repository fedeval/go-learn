package main

import "fmt"

func main() {
	// Declare a variable and its type explicitly
	var h int = 42
	fmt.Printf("The value of h is %d\n", h)

	// Declare with walrus (short declaration operator)
	// Go compiler figures out the type
	// NOTE: this can only be used inside a function
	a := 42
	fmt.Printf("The value of a is %d\n", a)

	// Can also do multiple assignments
	b, c, d := 1, "hello", 2.4
	fmt.Printf("b is %d, c is %s, d is %f\n", b, c, d)

	// When initialising a variable without an assigned
	// value the var gets the zero value for that given type
	var f int
	fmt.Println(f) // 0

	var g string
	fmt.Println(g) // ""

	var j struct{}
	fmt.Println(j) // nil

	/*
		It is not possible to assign a value of a certain type
		to a var of a different type, but it is possible to convert
		the type of the value to the right one in order to assign it.
	*/

	var p float32 = 3.4
	var s float64 = 4.4

	// p = s would throw an error, so we need to convert s
	p = float32(s)
	fmt.Printf("p has value %f and type %T\n", p, p) // 4.4, float32

}
