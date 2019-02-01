package main

import (
	"fmt"
)

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5}
	slice := arr[4:5]
	fmt.Println(slice, len(slice), cap(slice))

	slice1 := append(slice, 5, 6)
	fmt.Println(slice1, len(slice1), cap(slice1))

	slice2 := make([]int, len(slice1), cap(slice1)*2)

	copy(slice2, slice1)
	fmt.Println(slice2, len(slice2), cap(slice2))
}
