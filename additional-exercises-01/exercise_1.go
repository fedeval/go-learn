package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Age  int
}

func Exercise1() {
	u1 := user{
		Name: "James",
		Age:  32,
	}
	u2 := user{
		Name: "Jennie",
		Age:  27,
	}
	users := []user{u1, u2}
	bstring, err := json.Marshal(users)
	if err != nil {
		print(err)
	}
	fmt.Println(string(bstring))
}
