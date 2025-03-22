package main

import "fmt"

func main() {
	str := "hello你好"
	count := len([]rune(str))
	fmt.Printf("字符串'%s'的字符个数是:%d\n", str, count)
}
