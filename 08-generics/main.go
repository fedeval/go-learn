package main

import "fmt"

// The only difference between the two funcs below is the
// types they accept.
func AddI(x, y int) int {
	return x + y
}

func AddF(x, y float64) float64 {
	return x + y
}

// Generics allow us to define a similar
// func that would accept multiple types
func AddG[T int | float64](x, y T) T {
	return x + y
}

// We can further abstract this by using a
// type set (instead of a method set) to define
// an interface
type intOrFloat interface {
	int | float64
}

func AddG2[T intOrFloat](x, y T) T {
	return x + y
}

// Note that the interface above would exclude aliases
// with underlying type int or float64
type intAlias int

// We can use ~ to inclue aliases in the type set
type intOrFloatWithAliases interface {
	~int | ~float64
}

func AddG3[T intOrFloatWithAliases](x, y T) T {
	return x + y
}

// more at https://go.dev/doc/tutorial/generics
func main() {
	fmt.Println(AddI(1, 3))
	fmt.Println(AddF(1.2, 3.1))

	// Instead of using two differnt funcs we
	// can call the same func for two different types
	fmt.Println(AddG(1, 3))
	fmt.Println(AddG(1.2, 3.1))
	fmt.Println(AddG2(1, 3))
	fmt.Println(AddG2(1.2, 3.1))

	var z intAlias = 2

	// fmt.Println(AddG(z, 3)) both these would upset the compiler
	// fmt.Println(AddG2(z, 3))
	fmt.Println(AddG3(z, 7))

}
