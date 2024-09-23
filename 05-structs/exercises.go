package main

import "fmt"

type kid struct {
	firstName         string
	lastName          string
	favoriteIceCreams []string
}

func printInfo(k kid) {
	fmt.Printf("My name is %s %s and my favorite ice cream flavors are: \n", k.firstName, k.lastName)

	for _, v := range k.favoriteIceCreams {
		fmt.Println(v)
	}

}

func Exercise1() {
	jack := kid{
		firstName:         "Jackie",
		lastName:          "Stewart",
		favoriteIceCreams: []string{"Strawberry", "Vanilla"},
	}

	alan := kid{
		firstName:         "Alan",
		lastName:          "Jones",
		favoriteIceCreams: []string{"Chocolate", "Banana"},
	}

	kids := map[string]kid{
		jack.lastName: jack,
		alan.lastName: alan,
	}

	for _, kid := range kids {
		printInfo(kid)
	}
}

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	brand string
	model string
	doors int
	color string
}

func Exercise2() {
	v := vehicle{
		engine: engine{
			electric: false,
		},
		brand: "Ferrari",
		model: "F40",
		doors: 2,
		color: "Rosso Corsa",
	}
	fmt.Println(v)
	fmt.Printf("%T\n", v)
	fmt.Println(v.electric)
}
