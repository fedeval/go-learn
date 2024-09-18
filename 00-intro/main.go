package main

import "fmt"

func main() {
	fmt.Println("Hello, Gophers 😃")

	// Go has real immutable constants unlike Python
	// these values cannot be re-assigned.
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("%s is %d years old and age is of type %T while name is of type %T .\n", name, age, age, name)

	fmt.Printf(`
	You can write a multi-line
	string literal like this
	if you want to.
	`)
}
