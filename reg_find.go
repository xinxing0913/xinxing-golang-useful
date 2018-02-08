package main 

import (
    "fmt"
    "regexp"
)

func main() {
	// 首先编译一个正则表达式
	r, _ := regexp.Compile("[0-9]+")

	// 定义字符串
	str := "在2008年的时候，我8岁了，身高75厘米，上2年级"

	// 不限制提取次数
	arr1 := r.FindAllString(str, -1)
	fmt.Println("总的提取的数据:", arr1)
	fmt.Println("提取的第二个数据:", arr1[1])

    // 最多提取两次
	arr2 := r.FindAllString(str, 2)
	fmt.Println("总的提取的数据:", arr2)
	fmt.Println("提取的第二个数据:", arr2[1])
}