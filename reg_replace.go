package main 

import (
    "fmt"
    "regexp"
)

func main() {
	// 定义字符串
	str := "我今年22岁，学习java3年了，学习go2年了"

	// 定义正则表达式
	r, _ := regexp.Compile("[0-9]+")

	// 替换后的结果
	result := r.ReplaceAllString(str, "--保密--")
	fmt.Println("原始内容: ", str)
	fmt.Println("后来内容: ", result)
}