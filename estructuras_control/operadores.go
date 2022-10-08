package main

import "fmt"

func main() {
	//mayorEdad(-1)
	//edad(15)
	//menu("2")
	//bucle()
	//bucle2()
	menu()
}

func mayorEdad(a int) {
	if a > 0 && a < 18 {
		fmt.Println("Eres niño")
	} else if a >= 18 && a <= 65 {
		fmt.Println("Eres adulto")
	} else if a > 65 {
		fmt.Println("Eres anciano")
	} else {
		fmt.Println("Edad invalida!")
	}
}

func edad(a int) {
	switch {
	case a > 0 && a < 18:
		fmt.Println("Eres menor de edad!")
	case a >= 18 && a <= 65:
		fmt.Println("Eres mayor de edad!")
	case a > 65:
		fmt.Println("Eres un anciano!")
	default:
		fmt.Println("Edad no valida!")
	}
}

func bucle() {
	var iteraciones int
	fmt.Println("Ingrese el nro de iteraciones:")
	fmt.Scanln(&iteraciones)

	for i := 1; i <= iteraciones; i++ {
		fmt.Println("Iteracion nro ", i)
	}
}

func bucle2() {
	var array [10]int

	for i := 0; i < len(array); i++ {
		array[i] = i + 1
	}
	fmt.Println(array)
}
