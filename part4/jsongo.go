package main

import (
	"encoding/json"
	"fmt"
)

type animal struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	var bird animal
	JsonMy := `{"name": "Pigeon","description": "Nice but not flying but he is super"}`
	json.Unmarshal([]byte(JsonMy),&bird)
	fmt.Println(bird)
}
