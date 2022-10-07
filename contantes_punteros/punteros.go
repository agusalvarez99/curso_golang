package main

import "fmt"

func punteros() {
	var nombre string = "Agustin"
	var apellido string = "Alvarez"

	fmt.Println("Nombre: ", nombre, "Ubicacion: ", &nombre)
	fmt.Println("Apellido: ", apellido, "Ubicacion: ", &apellido)
}
