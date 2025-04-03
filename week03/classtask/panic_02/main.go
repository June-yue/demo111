package main

import "fmt"

func accessArray(a []int, index int) int {
	if index >= len(a) || index < 0 {
		panic("索引超出数组范围")
	}
	return a[index]
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v\n", r)
		}
	}()
	num := accessArray(a, 6)
	fmt.Printf("%d", num)
}
