package main

import "fmt"

func main(){
	channel := make(chan int)
	go factorial(5,channel)
	fmt.Println(<-channel)
	fmt.Println("the end")
}
func factorial(n int, ch chan int){
     
    result := 1
    for i := 1; i <= n; i++{
        result *= i
    }
    fmt.Println(n, "-", result)
	ch <- result
}