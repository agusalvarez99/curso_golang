package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	array := make([]int, n)
	fmt.Scan(&array[0], &array[1], &array[2], &array[3], &array[4])
	sort.Ints(array)

	if n%2 == 0 {
		fmt.Println((array[n/2-1] + array[n/2]) / 2)
	} else {
		fmt.Println(array[n/2])
	}
}
