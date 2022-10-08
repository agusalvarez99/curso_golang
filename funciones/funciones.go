package main

import "fmt"

func main() {

	/*suma(8, 2)
	res := retornaAlgo(10)
	fmt.Println("El precio sera de", res)
	var imc float64
	var cadena string
	imc, cadena = indice(60, 1.73)
	fmt.Println(imc)
	fmt.Println(cadena)
	parametrosVariables()
	parametrosVariables(5)
	parametrosVariables(5, 32, 55)
	fmt.Println(fibonacci(25))*/
	area := cuadrado(5)
	fmt.Println("El area es:", area())
}

func suma(n1 int, n2 int) {
	fmt.Println("La suma de ambos numeros es:", n1+n2)
}

func retornaAlgo(precio float32) float32 {
	return precio * 1.75
}

//otra forma de retornar un valor
func retornaAlgo2(precio float32) (final float32) {
	final = precio * 1.75
	return
}

func indice(peso float64, altura float64) (imc float64, literal string) {
	imc = peso / (altura * altura)
	literal = "Tu indice de masa corporal es de " + fmt.Sprintf("%f", imc)
	return imc, literal
}

//puede recibir cualquier cantidad de parametros
func parametrosVariables(nums ...int) {
	fmt.Println(nums, "")
}

//recursividad
func fibonacci(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return fibonacci(num-2) + fibonacci(num-1)
}

//funcion lambda
func cuadrado(lado float32) (area func() float32) {
	area = func() float32 {
		return lado * lado
	}
	return
}
