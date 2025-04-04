package main

import (
	"fmt"
	"sync"
)

func producer(num int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < num; i++ {
		ch <- i
		fmt.Printf("生产者发送:%d\n", i)
	}
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("消费者接收:%d\n", num)
	}
}

func main() {
	num := 10
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go producer(num, ch, &wg)
	wg.Add(1)
	go consumer(ch, &wg)
	wg.Wait()
}
