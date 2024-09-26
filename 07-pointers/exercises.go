package main

import "fmt"

func Exercise1() {
	a := "hello"
	fmt.Println(&a)
}

type person struct {
	name string
}

func changeNameVal(p person, newName string) person {
	p.name = newName
	return p
}

func changeNamePtr(p *person, newName string) {
	p.name = newName
}

func Exercise2() {
	p1 := person{"John"}
	changeNameVal(p1, "Jack")
	fmt.Println(p1)
	p1 = changeNameVal(p1, "Jack")
	fmt.Println(p1)

	p2 := person{"Mary"}
	changeNamePtr(&p2, "Rose")
	fmt.Println(p2)
}
