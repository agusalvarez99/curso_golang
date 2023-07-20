package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	contador := make(map[int]int)

	for _, nro := range a {
		contador[nro]++
	}

	var unique int

	for key, value := range contador {
		if value == 1 {
			unique = key
			break
		}
	}
	fmt.Println(unique)
}
