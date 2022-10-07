package main

import "fmt"

func main() {
	var animal string = "Gato"
	var cadena1, cadena2 string = "hola1", "hola2"

	fmt.Println(animal, cadena1, cadena2)

	var (
		cadena3 string = "hola"
		num     int    = 3
		flag    bool   = false
	)
	fmt.Println(num, cadena3, flag)

	var palabra string
	var entero int
	var flotante float32
	var booleano bool

	fmt.Println(palabra, entero, flotante, booleano)
	//la sintaxis que esta a continuacion sirve en el ambito de una funcion
	perro := "Miru"
	fmt.Printf("%v, %T\n", perro, perro)

}
