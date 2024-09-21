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
}
