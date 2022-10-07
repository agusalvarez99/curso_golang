package main

import "fmt"

const pi = 3.14159

func main() {
	const (
	//max = 500
	//min = 100
	)
	//punteros()
	//fmt.Println("El valor de PI es: ", pi)
	//fmt.Println("El valor max es: ", max)
	//fmt.Println("El valor min es: ", min)
	var precio float32 = 15
	fmt.Println(precio, "---", &precio)
	//precio = multiplicar(precio)
	multiplicarPuntero(&precio)
	fmt.Println(precio, "---", &precio)

}
