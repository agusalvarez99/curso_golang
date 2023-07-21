package main

import (
	"fmt"
)

func main() {
	var arr []int32
	// para probarlo, definir un array
	arr_sort := make([]int32, 100)

	for _, num := range arr {
		arr_sort[num]++
	}
	fmt.Println(arr_sort)
}
