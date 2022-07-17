package main

import "fmt"

type Animal interface{
	sound()
}
type cat struct{}
func (c *cat)sound(){
	fmt.Println("Mew")
}
type dog struct{}
func (d *dog) sound(){
	fmt.Println("Bark")
}
func main(){
	var Max Animal = &cat{}
	var Iggie Animal = &dog{}

	Max.sound()
	Iggie.sound()
}
