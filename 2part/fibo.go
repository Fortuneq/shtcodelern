package main

import "fmt"

func main() {
	ch := make(chan int)

	go factorial(5, ch)

	fmt.Println(<-ch)
}

func factorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
}
