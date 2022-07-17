package main

import "fmt"

type human struct{
	name string 
	age int
}

func(h *human)printall(){
	fmt.Println(h.name)
	fmt.Println(h.age)
}

func(h *human)eating(buter string){
	fmt.Printf("Human %s eating %s",h.name,buter)
}
func main(){
	var Alice human = human{"ALice",18}
	Alice.printall()
	Alice.eating("Sandwitch")
}