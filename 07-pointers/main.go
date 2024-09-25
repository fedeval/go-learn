package main

import "fmt"

func changeInt(n *int) {
	*n = 20
}

func changeIntWontWork(n int) {
	n = 30
}

func changeSlice(s []int) {
	s[0] = 2
}

func appendToSlice(s []int) {
	s = append(s, 1, 2, 3)
	s[0] = 5
}

func appendToSlicePtr(s *[]int) {
	*s = append(*s, 1, 2, 3)
}

type car struct {
	model string
}

func (c car) startEngine() {
	fmt.Printf("Starting %s\n", c.model)
}

func (c *car) shutDown() {
	fmt.Printf("Turning off %s\n", c.model)
}

type vehicle interface {
	startEngine()
	shutDown()
}

func startVehicle(v vehicle) {
	v.startEngine()
}

func main() {
	x := 42
	fmt.Println(x)  // prints the value of the variable
	fmt.Println(&x) // print the memory address at which the variable is stored

	// & is the address operator, it generates a pointer to the object
	fmt.Printf("%T\n", &x) // *int

	// * can be used to de-reference a pointer, that is, getting the value
	// of the object from its memory address
	xAddr := &x
	fmt.Println(*xAddr) // 42

	// we can use that to modify a value when we know its memory address
	*xAddr = 50
	fmt.Println(x) // 50

	/*
		In go functions, arguments are pass by value, meaning the value
		is copied and assigned to the corresponding argument when the function is called.
		So when passing a variable to a function, it won't be possible to modify the
		original value. However, that can be changed by passing a pointer to that
		variable instead (i.e. its memory address) and using de-referncing to modify it
	*/
	n := 10
	changeIntWontWork(n)
	fmt.Println(n) // 10
	changeInt(&n)
	fmt.Println(n) // 20

	/*
		IMPORTANT: composite types are reference types that have an underlying data
		structure to which they point. For example, a slice is a reference to an underlying
		array. So mutating a slice passed as a function argument will mutate the underlying
		array and thus the value of the slice outside of the function. Reference types
		behave like pointers.
	*/
	s := []int{0, 1, 2}
	changeSlice(s)
	fmt.Println(s) // [2,1,2]

	/*
		This does not apply to cases in which the mutation creates a new underlying array
		For instance if we append values above a slice capacity a new array is assigned to be
		the underlying data structure to the slice returned by the append, but that's only
		within the scope of the function and the original variable still points to the same
		underlying array so any modification done within the scope of the function won't
		change the original slice.
	*/
	appendToSlice(s)
	fmt.Println(s) // [2,1,2]

	// To make sure we modify the original slice even when
	// the underlying array changes we can make sure we
	// pass a pointer to the slice
	appendToSlicePtr(&s)
	fmt.Println(s) // [2,1,2,1,2,3]

	// When a method is added to a type, it also becomes part of the
	// available to the pointer to that type and viceversa
	fiat := car{"Fiat 500"}
	alfa := &car{"Alfa Romeo Giulietta"}

	fiat.startEngine()
	fiat.shutDown()
	alfa.startEngine()
	alfa.shutDown()

	// however, only the pointer type will have both methods as part
	// of its method set. This can be shown with interfaces.

	// the below would throw a compiler error since shutDown() is not
	// part of the cat method set so car does not fit the vehicle interface
	// startVehicle(fiat)

	// this works because both startEngine and shutDown are part of the method
	// sets of the *car pointer type so *car fits the vehicle interface
	startVehicle(alfa)
}
