package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}
	combinedSlice := append(slice1, slice2...)
	fmt.Println(combinedSlice)
	for i := 0; i < len(slice2); i++ {
		flag := true
		for j := 0; j < len(slice1); j++ {
			if slice1[j] == slice2[i] {
				flag = false
				break
			}
		}
		if flag {
			slice1 = append(slice1, slice2[i])
		}
	}
	fmt.Println(slice1)

}
