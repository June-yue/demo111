package main

import "fmt"

func divide(a int, b int) int {
	if b == 0 {
		panic("除数为零")
	}
	num := a / b
	return num
}

func main() {
	var a, b int
	fmt.Println("输入两个整数:")
	fmt.Scanln(&a, &b)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from:%v\n", r)
		}
	}()
	num := divide(a, b)
	fmt.Printf("%d", num)
}
