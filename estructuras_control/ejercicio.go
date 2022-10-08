package main

import "fmt"

func menu() {
	var opcion int
	fmt.Println("Ingrese la opcion que desee: ")
	fmt.Scanln(&opcion)
	switch opcion {
	case 1:
		recorrerArray()
	case 2:
		esPrimo()
	case 3:
		break
	default:
		fmt.Println("Opcion no valida!")
	}
}

func esPrimo() {
	var n, divisores int
	fmt.Println("Ingrese el nro a chequear!")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisores += 1
		}
	}

	if divisores == 2 {
		fmt.Print("El numero es primo!")
	} else {
		fmt.Print("El numero no es primo!")
	}
}

func recorrerArray() {
	num := 0
	var array [10]int

	for i := 0; i < len(array); i++ {
		array[i] = num + 1
		num += 1

	}
	fmt.Println(array)
}
