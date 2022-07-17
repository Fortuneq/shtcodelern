package main

import "fmt"

func main(){
	ch:= make(chan int,3)
	ch <- 123
	ch <- 19
	close(ch)

	for i:= 0;i < cap(ch);i++{
		if val,opened := <-ch;opened{
			fmt.Println(val)
		} else {
			fmt.Println("channel closed")
		}
	}
}