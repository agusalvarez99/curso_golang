package main

import "fmt"

func multiplicar(num float32) float32 {
	fmt.Println(num, "---", &num)
	return num * 10
}

func multiplicarPuntero(num *float32) {
	fmt.Println(num, "---", &num)
	*num = *num * 10

}
