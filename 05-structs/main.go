package main

import "fmt"

type person struct {
	name string
	age  int
}

// a struct can be a property of another struct
type doctor struct {
	person
	field string
}

func main() {
	// sample struct
	j := person{
		name: "john",
		age:  11,
	}
	fmt.Printf("%T\n", j)
	fmt.Println(j)
	fmt.Println(j.name) // john

	// struct with embedded struct
	dr := doctor{
		person: person{
			name: "mike",
			age:  40,
		},
		field: "cardiology",
	}

	fmt.Println(dr)

	// properties of the embedded struct can also be accessed directly
	fmt.Println(dr.person.name) // mike
	fmt.Println(dr.name)        // mike

	// structs can be created anonymously as well
	m := struct {
		name string
		age  int
	}{
		name: "Test",
		age:  20,
	}

	fmt.Printf("%T\n", m)

	// anonymous types are indeterminate and can be assigned to
	// a variable of a given type if the shape matches, the compiler
	// will infer the type
	var p person = m
	fmt.Printf("%T\n", p)

	// exercises
	Exercise1()
	Exercise2()
}
