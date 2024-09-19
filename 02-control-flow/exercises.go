package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("Print this first")
}

func Exercise1() {
	n250 := rand.Intn(251)
	fmt.Println("The random int is", n250)

	if n250 <= 100 {
		fmt.Println("Between 0 and 100")
	} else if n250 <= 200 {
		fmt.Println("Between 101 and 200")
	} else {
		fmt.Println("Between 201 and 250")
	}
}

func Exercise2() {
	n250 := rand.Intn(251)

	switch {
	case n250 <= 100:
		fmt.Println("Between 0 and 100")
	case n250 > 100 && n250 <= 200:
		fmt.Println("Between 101 and 200")
	case n250 > 200 && n250 <= 250:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("This will never print.")
	}

}

func Exercise3() {
	for i := 0; i < 100; i++ {
		fmt.Println("i is: ", i)
	}

	n := 0
	for n < 100 {
		fmt.Println("n is: ", n)
		n++
	}

	x := 0
	for {
		fmt.Println("x is: ", x)
		if x == 99 {
			break
		}
		x++
	}

}

func Exercise4() {

	for n := 0; n < 20; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Printf("%v is an odd number\n", n)
	}
}

func Exercise5() {
	people := map[string]int{
		"John": 20,
		"Mary": 21,
		"Jack": 22,
	}

	for name, age := range people {
		fmt.Printf("%s is %v years old\n", name, age)
	}
}

func Exercise6() {
	people := map[string]int{
		"John": 20,
		"Mary": 21,
		"Jack": 22,
	}

	valuesToCheck := []string{"John", "Marco"}

	for _, v := range valuesToCheck {
		if age, exists := people[v]; exists {
			fmt.Printf("%s exists and is %v years old\n", v, age)
			continue
		}
		fmt.Printf("%s does not exist\n", v)
	}
}
