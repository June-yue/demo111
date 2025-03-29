package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var count int = 0

func rule() {
	fmt.Println("欢迎来到猜字数游戏!")
	fmt.Println("规则:")
	fmt.Println("1.计算机将在1到100之间随机选择一个数字。")
	fmt.Println("2.你可以选择难度级别（简单、中等、困难)，不同难度对应不同的猜测机会。")
	fmt.Println("3.请输入您的猜测。")
	fmt.Println()
	fmt.Println("请选择难度级别（简单/中等/困难):")
	fmt.Println("1.简单(10次机会)")
	fmt.Println("2.中等(5次机会)")
	fmt.Println("3.困难(3次机会)")
	fmt.Println()
}

func random() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1
	return num

}

func guess(target, number int, start time.Time) {
	var num int
	fmt.Println()
	for i := 1; i <= number; i++ {
		fmt.Printf("第%d次猜测,请输入您的数字(1-100):", i)
		fmt.Scanln(&num)
		if num > target {
			fmt.Printf("您猜的数字大了。\n")
		} else if num < target {
			fmt.Printf("您猜的数字小了。\n")
		} else {
			fmt.Printf("恭喜您猜对了！您在第%d次猜测中成功。\n", i)
			now := time.Now()
			layout := "2006-01-02 15:04:05"
			writegame("游戏结束时间：" + now.Format(layout))
			duration := now.Sub(start)
			writegame("所用时间：" + duration.String())
			writegame("游戏结果：成功")
			writegame(" ")
			again()
			return
		}
		writegame("第一次猜测数字为：" + strconv.Itoa(num))
	}
	now := time.Now()
	layout := "2006-01-02 15:04:05"
	writegame("游戏结束时间：" + now.Format(layout))
	duration := now.Sub(start)
	writegame("所用时间：" + duration.String())
	writegame("游戏结果：失败")
	writegame(" ")
	fmt.Println("很遗憾，机会已经用完了。\n")
	again()
}

func again() {
	fmt.Println("是否重新开始游戏？")
	var choice int
	fmt.Println("1.重新开始游戏\n2.退出游戏")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		game()
	case 2:
		return
	}
}

func game() {
	target := random()
	now := time.Now()
	writegame("随机数为：" + strconv.Itoa(target))
	layout := "2006-01-02 15:04:05"
	var num int
	count++
	writegame("游戏次数：" + strconv.Itoa(count))
	writegame("游戏开始时间：" + now.Format(layout))
	fmt.Printf("输入选择：")
	fmt.Scanln(&num)
	fmt.Printf("开始游戏：")
	fmt.Println()
	switch num {
	case 1:
		guess(target, 10, now)
	case 2:
		guess(target, 5, now)
	case 3:
		guess(target, 3, now)
	default:
		fmt.Printf("错误")
	}
}

func writegame(str string) {
	file, err := os.OpenFile("./game.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,err:%v", err)
		return
	}
	file.WriteString(str + "\n")
	file.Close()
}

func main() {
	rule()
	game()
}
