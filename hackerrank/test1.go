package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	//validar que sea impar
	if n%2 == 0 {
		fmt.Println("El numero ingresado debe ser impar!")
	} else {
		array := make([]int, n)
		//no se cuantos elementos son
		for i := 0; i < n; i++ {
			fmt.Scan(&array[i])
		}
		sort.Ints(array)
		median := array[n/2]
		fmt.Println(median)
	}
}
