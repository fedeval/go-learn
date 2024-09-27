package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func Exercise2() {
	s := `[{"Name":"James","Age":32},{"Name":"Jennie","Age":27}]`
	var people []person

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}
	for _, p := range people {
		fmt.Println("My name is", p.Name, "and I am", p.Age, "years old.")
	}
}
