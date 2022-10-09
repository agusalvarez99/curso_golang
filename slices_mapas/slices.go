package main

import "fmt"

func main() {

	/*var animales = make([]string, 3)
	animales[0] = "gato"
	animales[1] = "perro"
	animales[2] = "caballo"

	animales = append(animales, "conejo")
	fmt.Println(animales)

	//borrar elemento
	animales = append([]string{}, animales[0], animales[1], animales[3])
	fmt.Println(animales)

	animales2 := animales[0:2]
	fmt.Println(animales2)

	animales3 := animales[2:]
	fmt.Println(animales3)

	animales = append(animales2, animales3...)
	fmt.Println(animales)*/
	//Copiar()
	mapa()
}

func Copiar() {
	var personas = make([]string, 2)
	personas[0] = "Juan"
	personas[1] = "Agustin"
	personas = append(personas, "Marcos", "Pedro")
	fmt.Println(personas)

	//no es una copia, apunta a la misma direccion de memoria
	/*personas2 := personas
	fmt.Println(personas2)
	personas[0] = "Maria"

	fmt.Println(personas2, personas)*/

	//esto si realiza una copia, tambien existe una funcion copy()
	personas2 := append([]string{}, personas...)
	fmt.Println(personas2)
}
