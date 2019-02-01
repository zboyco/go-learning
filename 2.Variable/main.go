package main

import (
	"fmt"
)

// 普通声明
var a int = 10

// 自行判定变量类型
var b = 10

// 多个变量
var x, y int

// 因式分解
var (
	c string
	d bool
)

func main() {
	//这种不带声明格式的只能在函数体中出现
	c := 10
	v1, v2, v3 := 1, 2, 3
	fmt.Println(a, b, c, v1, v2, v3, c, d)
}
