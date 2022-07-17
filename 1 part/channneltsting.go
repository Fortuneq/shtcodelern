package main

import "fmt"

func main(){
	fmt.Println("Start")
	fmt.Println(<-createchannel(10))
	fmt.Println("succes")
}

func createchannel(n int)chan int{
	ch := make(chan int)
	go func(){
		ch <-n
	}()
	return ch
}