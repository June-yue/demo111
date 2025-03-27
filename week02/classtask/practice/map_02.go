package main

import "fmt"

func main() {
	str := "hello world123"
	count := make(map[rune]int)
	for _, char := range str {
		count[char]++
	}
	max := 0
	var c rune
	for char, num := range count {
		if num > max {
			max = num
			c = char
		}
	}
	fmt.Printf("出现次数最多的字母为：%c", c)
}
