package main

import "fmt"

/*
Go supports polymorphism, meaning a value can be of multiple types.
This is done through interfaces.
A type satisfies the requirements of an interface if it implements
all the methods specified by the interface.
Note: a type can implement any additional methods and it would still
be satisfying an interface as long as the above mentioned condition
is met.
*/
type Dog struct {
	name  string
	breed string
	sound string
}

type Cat struct {
	name     string
	furColor string
	sound    string
}

type Bird struct {
	name  string
	sound string
}

func (d Dog) sayHello() {
	fmt.Printf("The %s called %s says %s!\n", d.breed, d.name, d.sound)
}

func (c Cat) sayHello() {
	fmt.Printf("The %s cat called %s says %s!\n", c.furColor, c.name, c.sound)
}

func (b Bird) sayHello() {
	fmt.Printf("The bird called %s says %s!\n", b.name, b.sound)
}

type Animal interface {
	sayHello()
}

func makeAnimalSayHi(a Animal) {
	a.sayHello()
}

func testInterfaces() {
	spike := Dog{
		name:  "Spike",
		breed: "Labrador",
		sound: "Woof",
	}

	timmy := Cat{
		name:     "Timmy",
		furColor: "orange",
		sound:    "Meow",
	}

	johnny := Bird{
		name:  "Johnny",
		sound: "Chirp Chirp",
	}

	/*
		johnny, spike and timmy are Bird, Dog and Cat type, but those types implement
		the sayHello method which means they satisfy the Animal inteface,
		hence they can be assigned as values in a slice of Animal type values.
	*/
	animals := []Animal{spike, timmy, johnny}

	for _, animal := range animals {
		/*
			since we know Bird, Dog and Cat both implement the Animal interface, we can
			call safely pass them to a function that takes an Animal value and calls
			the sayHello method on it
		*/
		makeAnimalSayHi(animal)
	}
}
