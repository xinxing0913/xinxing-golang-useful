package main 

import (
    "fmt"
    "regexp"
)

func main() {
	// 直接验证某个字符串是否符合某个模式
	match, _ := regexp.MatchString("1[0-9]{10}", "15722223333")
	fmt.Println("匹配结果是:", match)

	// 首先进行编译
	r, _ := regexp.Compile("x[a-z]?")

	// 然后判断是否匹配xinxing
	match1 := r.MatchString("xinxing")
	fmt.Println("是否匹配xinxing:", match1)

	// 然后判断是否匹配java
	match2 := r.MatchString("java")
	fmt.Println("是否匹配java:", match2)

}