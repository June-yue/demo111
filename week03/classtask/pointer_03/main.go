package main

import "fmt"

func doublesValues(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] = (*nums)[i] * 2
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	doublesValues(&nums)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i])
	}

}
