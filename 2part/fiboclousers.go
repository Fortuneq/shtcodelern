package main

import "fmt"

func main() {
	results := make(map[int]int)
	structch := make(chan struct{})

	go factorial(5, structch, results) //запускаем горутину с факториалом
	//и по идее она должна даже не успеть выполнится но мы на следующей строке читаем инфу из канала, а  в него пока ничего не записали

	<-structch //читаем инфу после defer поэтому канал закрыт инфа прочитана в никуда продолжаем наш тур де франс

	for i, v := range results {
		fmt.Println(i, "-", v)
	}

}

func factorial(n int, ch chan struct{}, results map[int]int) {
	defer close(ch) // закрываем канал после завершения горутины
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		results[i] = result
	}
}
