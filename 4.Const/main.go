package main

import (
	"fmt"
)

const name1 string = "baiyang"
const name2 = "mia"

// 常量用作枚举
const (
	Unknow = 0
	Female = 1
	Male   = 2
)

// iota 特殊常量
// 赋值不写，默认为上一个表达式
const (
	a = iota //0
	b        //1
	c        //2
	d = "ha" //独立值，iota += 1
	e        //"ha"   iota += 1
	f = 100  //iota +=1
	g        //100  iota +=1
	h = iota //7,恢复计数
	i        //8
)

func main() {
	fmt.Println(name1, name2)
	fmt.Println(Unknow, Female, Male)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
