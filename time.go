package main 

import (
    "fmt"
    "time"
)


func main() {
	func1()
	func2()
	func3()
	func4()
}

// 转换为时间戳
func func1() {
	// 获取当前时间的时间戳
	fmt.Println("当前时间的时间戳:", time.Now().Unix())
	// 获取特定时间的时间戳
	my_time, _ := time.Parse("2006-01-02 15:04:05", "2008-08-08 20:00:00")
	fmt.Println("北京奥运会开始的时间戳:", my_time.Unix())
}

// 打印格式化的时间
func func2() {
	// 打印全格式
	fmt.Println("格式一:", time.Now().Format("2006-01-02 15:04:05"))
	// 只打印年月日
	fmt.Println("格式二:", time.Now().Format("2006-01-02"))
	// 只打印时分秒
	fmt.Println("格式三:", time.Now().Format("15:04:05"))
	// 打印年月日和小时
	fmt.Println("格式四:", time.Now().Format("2006-01-02 15"))
    // 用斜线的方式打印
    fmt.Println("格式五:", time.Now().Format("2006/01/02 15/04/05"))
}

// 输出当前星期几
func func3() {
	// 输出今天星期几
	fmt.Println("今天是:", time.Now().Weekday())
}

// 获取具体的时间单位
func func4() {
	now := time.Now()
	fmt.Println("当前的年份:", now.Year())
	fmt.Println("当前的月份:", now.Month())
	fmt.Println("当前的天数:", now.Day())
	fmt.Println("当前的小时:", now.Hour())
	fmt.Println("当前的分钟:", now.Minute())
	fmt.Println("当前的秒数:", now.Second())
}