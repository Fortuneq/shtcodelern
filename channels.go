package main

import "fmt"

func main(){
	intchan := make(chan int)
	go func(){
		fmt.Println("Programm write in channel")
		intchan<-5
	}()
	fmt.Println(<-intchan)
	fmt.Println("end")
}