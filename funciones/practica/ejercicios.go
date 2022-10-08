package main

import "fmt"

func main() {
	/*var producto float32
	var suma float32
	var division float32
	producto, suma, division = math(5, 10)
	fmt.Println("Multiplicacion: ", producto)
	fmt.Println("Suma: ", suma)
	fmt.Println("Division: ", division)
	fmt.Println("El factorial es:", factorial(5))*/
	activ2()
}

func math(num1 float32, num2 float32) (mul float32, sum float32, div float32) {
	mul = num1 * num2
	sum = num1 + num2
	div = num1 / num2
	return
}

func factorial(num float32) float32 {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

type CV struct {
	nombre      string
	email       string
	calle       string
	nroCalle    int
	experiencia [2]Experiencia
}

type Experiencia struct {
	tecnologia string
	aniosExp   int
}

func (cv *CV) mostrarDatos() {
	fmt.Println("Nombre:", cv.nombre)
	fmt.Println("Email:", cv.email)
	fmt.Println("Direccion:", cv.calle, cv.nroCalle)
	fmt.Println("Experiencia:", cv.experiencia)
}

func activ2() {
	var exp [2]Experiencia
	exp[0].tecnologia = "Python"
	exp[0].aniosExp = 5
	exp[1].tecnologia = "GO"
	exp[1].aniosExp = 1

	cv := CV{
		"agustin",
		"agus@mail.com",
		"belgrano",
		120,
		exp,
	}
	cv.mostrarDatos()
}
