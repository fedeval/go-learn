package main

import "fmt"

func main() {
	// Map literal
	a := map[string]int{
		"john": 10,
		"mary": 12,
	}
	fmt.Println(a)

	// Like for slices, we can use make
	b := make(map[string]int)
	b["mark"] = 14
	fmt.Println(b)

	// maps also have a length propery
	fmt.Println(len(a)) // 2

	// we can also range over maps
	// keys only
	for name := range a {
		fmt.Println("My name is", name)
	}

	// keys and values
	for name, age := range a {
		fmt.Printf("My name is %v and I am %v years old\n", name, age)
	}

	// unlike slices, there is a built-in method to remove elements from maps
	delete(a, "john")
	fmt.Println(a)

	// note: go does not throw key errors, accessing an inexistent key
	// just returns 0
	j := a["john"]
	fmt.Println(j)

	// this can be confusing as there is no way to know if 0 was the actual value
	// in the map or the key did not exist. To counteract that we can use the the
	// comma ok idiom
	test := map[string]int{"jack": 0}

	jack, ok1 := test["jack"]
	john, ok2 := test["john"]

	fmt.Println(jack, ok1) // 0, true
	fmt.Println(john, ok2) // 0, false

	// Exercises
	Exercise1()
	Exercise2()
}
