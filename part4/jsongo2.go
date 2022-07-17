package main

import (
	"encoding/json"
	"fmt"
)

type Parameters struct {
	Heigth int `json:"height"`
	Weight int `json:"weight"`
}
type Figure struct {
	Name       string `json:"name"`
	Color      string `json:"color"`
	Parameters Parameters
}

func main() {
	myJson := `{"name":"square","color":"black","parameters": {"height":24,"weight":12}}`

	var square Figure

	json.Unmarshal([]byte(myJson), &square)

	fmt.Println(square)
}
