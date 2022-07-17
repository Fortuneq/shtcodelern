package main

import "fmt"

type Animal interface {
	speak()
}

type cat struct{ name string }

func (c *cat) speak() {
	fmt.Println(c.name, "MEWMEWEEWMEMW")
}

type dog struct{ name string }

func (d *dog) speak() {
	fmt.Println(d.name, "BARK")
}

type horse struct{ name string }

func (h *horse) speak() {
	fmt.Println(h.name, "IGOGO")
}

func main() {
	var cat Animal = &cat{"Alice"}

	var dog Animal = &dog{"Max"}

	var horse Animal = &horse{"Imperor"}

	animals := [...]Animal{cat, dog, horse}
	for _, animal := range animals {
		animal.speak()
	}
}
