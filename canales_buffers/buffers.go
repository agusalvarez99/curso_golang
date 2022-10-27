package main

import "fmt"

func EjemploBuffer() {
	//en el segundo parametro ponemos el tamaño que tendra el buffer
	canal := make(chan string, 1)
	//no podria almacenar otro valor en el buffer, xq el tamaño asignado fue de 1
	canal <- "Agustin"
	fmt.Println(<-canal)
	//en este punto ya esta liberado el buffer y ahi si podriamos
	//asignarle otro valor
	canal <- "Alvarez"
	fmt.Println(<-canal)

}

func EjemploBuffer2() {
	canal := make(chan string, 3)
	canal <- "Agustin"
	canal <- "Ezequiel"
	canal <- "Alvarez"
	close(canal)
	//iterar los valores que va tomando el canal
	for c := range canal {
		fmt.Println(c)
	}
}
