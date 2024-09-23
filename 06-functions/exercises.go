package main

import (
	"fmt"
	"math"
)

func Exercise1() {
	variadicFunc := func(s ...float32) (x float32) {
		for _, v := range s {
			x += v
		}
		return x
	}
	floats := []float32{1.2, 1.5, 5.5, 8.32}
	fmt.Println(variadicFunc(floats...))
}

func Exercise2() {
	// Note: defer is LIFO! (stack)
	// more here: https://go.dev/blog/defer-panic-and-recover
	defer fmt.Println("Slim Shady")
	defer fmt.Println("My name is, chka-chka")
	defer fmt.Println("My name is, who?")
	fmt.Println("Hi")
	fmt.Println("My name is, what?")
}

type Car struct {
	cv         int
	engineType string
	model      string
}

func (c Car) startEngine() {
	fmt.Printf("Started your %s, all the %v CVs of the %s are roaring!!\n", c.model, c.cv, c.engineType)
}

func Exercise3() {
	daytona := Car{
		cv:         840,
		model:      "Daytona SP3",
		engineType: "V12",
	}

	daytona.startEngine()
}

type Square struct {
	lenght float64
	width  float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.lenght * s.width
}

func (c Circle) Area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

type Shape interface {
	Area() float64
}

func info(s Shape) {
	fmt.Printf("The area of this shape is %v\n", s.Area())
}

func Exercise4() {
	c := Circle{
		radius: 2.54,
	}

	s := Square{
		lenght: 5.5,
		width:  6.4,
	}

	info(c)
	info(s)
}
