package main

import "fmt"

// function syntax
// func (r receiver) identifier (p parameters) (returns) { code }
func foo(s1, s2 string) string {
	return fmt.Sprintf("Hello %s and %s", s1, s2)
}

// functions can have variadic parameters (i.e. unlimited amount)
// an example of that is the built-in append function
// the variadic param will be a slice of the specified type
func sum(x ...int) (sum int) {
	fmt.Printf("Value: %v, Type: %T\n", x, x)
	for _, v := range x {
		sum += v
	}
	// this is a "naked" return, it can be done when the returning
	// value is named
	return
}

// the defer statement can be used to delay the execution of a function
// until the surrounding function returns
func worldHello() {
	// hello won't be printed until the worldHello function returns
	defer fmt.Println("hello")
	fmt.Println("world")
}

// we can create methods by attaching functions to a type
// by specifying the type as a receiver
type person struct {
	name string
}

func (p person) sayHello() {
	fmt.Printf("Hello, my name is %s\n", p.name)
}

func main() {
	fmt.Println(foo("Jack", "John"))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	// not passing any values for a variadic param does not result in an error
	// instead the param will be nil (the zero value of a slice)
	sum()

	worldHello()

	// methods are called by using dot notation on values of the type they are
	// attached to.
	j := person{
		name: "Jack",
	}
	j.sayHello()

	// it is possible to declare anonymous functions which are executed
	// right after being declared
	func() {
		fmt.Println("I am anonymous!")
	}()

	// anonymous functions can still take parameters
	func(s string) {
		fmt.Println("hello", s)
	}("John")

	// functions are first-class citizens in go so they can
	// be assigned to variables, be passed as params to
	// other function or be returned by other functions
	b := func() {
		fmt.Println("I am a first class citizen :)")
	}
	b()

	// see separate section on interfaces
	fmt.Println()
	fmt.Println("---------- Interfaces")
	testInterfaces()

	// exercises
	fmt.Println()
	fmt.Println("---------- Exercises")
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
}
