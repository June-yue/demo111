package main

import "fmt"

// 如果有函数，可以这里定义
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := slice[2:7]
	fmt.Println(slice1)
	slice1 = append(slice1, 100)
	fmt.Println(slice1)
	sum := 0
	for i := 0; i < len(slice1); i++ {
		sum += slice1[i]
	}
	fmt.Println(sum)
	//这里写代码
}
