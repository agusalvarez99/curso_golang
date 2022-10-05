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

func tipoArray() {
	//definicion de un array de tipo string con 4 elementos
	var nombres [4]string
	nombres[0] = "pedro"
	nombres[1] = "juan"
	nombres[2] = "meli"
	nombres[3] = "agus"
	fmt.Println(nombres[0], nombres[1], nombres[2], nombres[3])

	var numeros [3]int
	numeros[0] = 1
	numeros[1] = 2
	numeros[2] = 3
	fmt.Println(numeros[0], numeros[1], numeros[2])
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

type Persona struct {
	edad      int
	nombre    string
	apellido  string
	direccion Direccion
}

type Direccion struct {
	calle string
	nro   int
}

func (p *Persona) presentacion() {
	fmt.Printf("Hola soy %s %s y vivo en %s %d", p.nombre, p.apellido, p.direccion.calle, p.direccion.nro)
}

func tiposPropios() {

	dir := Direccion{
		"Belgrano",
		120,
	}
	persona := Persona{
		24,
		"Agustin",
		"Alvarez",
		dir,
	}
	//fmt.Println(persona)
	persona.presentacion()
}

func main() {
	//var txt string = "El resultado es "
	//fmt.Println(txt, suma(5, 6))
	//varBool()
	//println(tiposNum(5, 4.8))
	//tiposString("hola")
	//animal()
	//tipoArray()
	tiposPropios()
}
