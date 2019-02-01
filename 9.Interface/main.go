package main

import (
	"fmt"
)

type Phone interface {
	Call(name string) string
}

type Nokia struct {
}

func (phone Nokia) Call(name string) string {
	return "Nokia Say: Hello, " + name
}

type IPhone struct{}

func (phone IPhone) Call(name string) (result string) {
	// 方法的两种返回值方式
	// return "IPhone say : Hello, " + name
	result = "IPhone say : Hello, " + name
	return
}

func main() {
	var phone Phone
	phone = new(Nokia)
	fmt.Println(phone.Call("Tom"))

	phone = new(IPhone)
	fmt.Println(phone.Call("Jerry"))
}
