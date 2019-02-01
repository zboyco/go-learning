package main

import (
	"fmt"
)

func main() {
	b := 15

	number := []int{1, 2, 3, 5}

	for a := 0; a < b; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	var a int
	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}
	for key, value := range number {
		fmt.Printf("第 %d 位 x 的值 = %d\n", key, value)
	}
}
