package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Exercise3() {
	fmt.Println(Agents)
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(Agents)
}
