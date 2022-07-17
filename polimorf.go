package main

import "fmt"

type Animal interface{
	speak()
}

type cat struct {}

func(c *cat)speak(){
	fmt.Println("MEWMEWEEWMEMW")
}
type dog struct {}

func(d *dog)speak(){
	fmt.Println("BARK")
}
type horse struct{}

func(h *horse)speak(){
	fmt.Println("IGOGO")
}

func main(){
	var cat Animal = &cat{}

	var cat2 Animal = cat

	var dog Animal = &dog{}

	var horse Animal = &horse{}

	animals:= [...]Animal{cat,cat2,dog,horse}
	for _,animal:= range(animals){
		animal.speak()
	}
}