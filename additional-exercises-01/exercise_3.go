package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type agent struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func Exercise3() {
	u1 := agent{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := agent{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := agent{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	agents := []agent{u1, u2, u3}

	fmt.Println(agents)

	// your code goes here
	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode(agents)

}
