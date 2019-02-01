package main

import (
	"fmt"
)

func main() {
	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = i + 100
	}
	for i := 0; i < 10; i++ {
		fmt.Println(array[i])
	}
	for index, item := range array {
		fmt.Printf("第 %d 位的值是: %d\n", index, item)
	}
}
