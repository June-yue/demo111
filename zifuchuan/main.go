package main

import "fmt"

func main() {
	str := "china中国"
	count := len([]rune(str))
	fmt.Println("字符'%s'的个数是%d\n", str, count)

}
