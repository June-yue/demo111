package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := slice[2:7]
	slice1 = append(slice1, 11, 12, 13)
	slice1 = append(slice1[:4], slice1[5:]...)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = slice1[i] * 2
	}
	fmt.Println(slice1, cap(slice1))
}
