package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Example of if statement
	x := 9001
	if x < 9000 {
		fmt.Println("It's less than 9000")
	} else if x == 9000 {
		fmt.Println("It's exactly 9000")
	} else {
		fmt.Println("IT'S OVER 9000!!!!!!")
	}

	/*
		It's possible to execute statements within if statements, those
		statements are executed before the expression is evaluated.
		The scope of z in this case is limited to the if statement
	*/
	if z := int(900.2 * 10); z > x {
		fmt.Println("IT'S ALSO OVER 9000!!!!!!")
	}

	// Switch statements
	values := []int{8999, 9000, 90001, 0}
	switch values[rand.Intn(len(values))] {
	case 8999:
		fmt.Println("It's less than 9000")
	case 9000:
		fmt.Println("It's exactly 9000")
	case 9001:
		fmt.Println("IT'S OVER 9000!!!!!!")
	default:
		fmt.Println("Dunno")
	}

	switch {
	case x < 9000:
		fmt.Println("It's less than 9000")
	case x == 9000:
		fmt.Println("It's exactly 9000")
	case x > 9000:
		fmt.Println("IT'S OVER 9000!!!!!!")
	default:
		fmt.Println("Dunno")
	}

	// For loops
	// for init; condition; after
	for i := 1; i < 10; i++ {
		fmt.Println("i is", i)
	}

	// while-like for loop
	// can also be done with if + break
	a := 5
	for a < 10 {
		fmt.Println("a is", a)
		a++
	}

	// continue statement
	for l := 1; l < 10; l++ {
		if l%2 != 0 {
			// go to next iteration if condition is met
			// this is like Python
			continue
		}
		fmt.Println("l is an even number: ", l)
	}

	// we can iterate over indexed data structures
	// like slices or maps using range
	ints := []int{1, 2, 3, 4, 5}
	for i, v := range ints {
		fmt.Printf("Item with value %d at index %d\n", v, i)
	}

	// Excercises
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
}
