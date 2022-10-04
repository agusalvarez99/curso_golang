package main

import "fmt"

func suma(a int, b int) int {
	return a + b
}

func varBool() {
	var todoOK bool = true // por defecto es false
	fmt.Println(todoOK)
}

func main() {
	var txt string = "El resultado es "
	fmt.Println(txt, suma(5, 6))
	varBool()
}
