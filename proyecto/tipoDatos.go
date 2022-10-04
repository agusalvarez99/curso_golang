package main

import "fmt"

func suma(a int, b int) int {
	return a + b
}

func varBool() {
	var todoOK bool = true // por defecto es false
	fmt.Println(todoOK)
}

func tiposNum(entero int, decimal float32) float32 {
	//fmt.Printf("Entero: %v - Decimal: %v\n", entero, decimal)
	return (float32(entero) + decimal)
}
func tiposString(palabra string) {
	fmt.Println(palabra)
}

func animal() {
	var animal string
	var edad int

	fmt.Println("¿Cual es tu animal favorito?")
	fmt.Scanf("%s\n", &animal)

	fmt.Println("¿Que edad tiene?")
	fmt.Scanf("%d\n", &edad)

	fmt.Println("Mi animal favorito es el "+animal+" y tiene", edad, "años")
}

func main() {
	//var txt string = "El resultado es "
	//fmt.Println(txt, suma(5, 6))
	//varBool()
	//println(tiposNum(5, 4.8))
	//tiposString("hola")
	animal()
}
