package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	work :=func(id int){
		defer wg.Done()
		fmt.Println("Functon start to work")
		time.Sleep(2*time.Second)
		fmt.Println("End of gourotin func")
	}
	go work(1)
	go work(2)
	wg.Wait()
	fmt.Println("функции смогли выполнится ")

}

