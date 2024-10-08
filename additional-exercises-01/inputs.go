package main

type Agent struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

var a1 Agent = Agent{
	First: "James",
	Last:  "Bond",
	Age:   32,
	Sayings: []string{
		"Shaken, not stirred",
		"Youth is no guarantee of innovation",
		"In his majesty's royal service",
	},
}

var a2 Agent = Agent{
	First: "Miss",
	Last:  "Moneypenny",
	Age:   27,
	Sayings: []string{
		"James, it is soo good to see you",
		"Would you like me to take care of that for you, James?",
		"I would really prefer to be a secret agent myself.",
	},
}

var a3 = Agent{
	First: "M",
	Last:  "Hmmmm",
	Age:   54,
	Sayings: []string{
		"Oh, James. You didn't.",
		"Dear God, what has James done now?",
		"Can someone please tell me where James Bond is?",
	},
}

var Agents []Agent = []Agent{a1, a2, a3}
