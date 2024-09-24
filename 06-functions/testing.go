package main

func Add(x int, y int) int {
	return x + y
}

func Subtract(x int, y int) int {
	return x - y
}

func DoOperation(x int, y int, f func(int, int) int) int {
	return f(x, y)
}
