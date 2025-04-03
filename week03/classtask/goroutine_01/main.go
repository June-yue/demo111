package main

import "fmt"

func splitslice(s []int, ch1 chan int, ch2 chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- s[i] * s[i]
		}
		close(ch1)
	}()
	go func() {
		for i := 10; i < len(s); i++ {
			ch2 <- s[i] * s[i]
		}
		close(ch2)
	}()
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	ch1 := make(chan int)
	ch2 := make(chan int)
	splitslice(slice, ch1, ch2)
	num1 := 0
	num2 := 0
	for num := range ch1 {
		num1 += num
	}
	for num := range ch2 {
		num2 += num
	}
	fmt.Printf("%d", num1+num2)
}
