package main

import "fmt"

func main() {
	var str string
	fmt.Printf("请输入一个整数：")
	fmt.Scanf("%s", &str)
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-1-i] {
			fmt.Printf("false")
			return
		}
	}
	fmt.Printf("true")
}
