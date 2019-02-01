package main

import (
	"fmt"
)

var v1 int = 1
var a, b int
var c string

func main() {
	var ip *int
	// 使用 & 获取变量指针(内存地址)
	ip = &v1
	// 使用 * 访问指针(内存地址)
	fmt.Println(v1, ip, *ip)

	// 多变量可以在同一行进行赋值
	a, b, c = 5, 7, "abc"
	fmt.Println(a, b)
	// 为声明的变量不能使用
	// err = "err"

	// 可以直接交换变量
	a, b = b, a
	fmt.Println(a, b)

	// 标识符 _  符号可以抛弃值
	_, x := a, b
	fmt.Println(x)

	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
	/* ptr 不是空指针 */
	if ptr != nil {
		fmt.Println("不是空指针")
	}
	/* ptr 是空指针 */
	if ptr == nil {
		fmt.Println("是空指针")
	}

	// 指向指针变量的指针变量
	tr := 100
	var pptr **int
	// ptr指向tr的地址
	ptr = &tr
	// pptr指向ptr的地址
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 tr = %d\n", tr)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

}
