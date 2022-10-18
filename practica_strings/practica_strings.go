package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//textoMulti()
	//concatenarStrings()
	//convertir()
	//metodos()
	//busqueda()
	borrarEspacios()
}

func runes() {
	prueba := "Me llamo Agustin \t y mi apellido es \n Alvarez. \r Tengo 23 años"
	fmt.Println(prueba)

	array := []rune(prueba)
	fmt.Println(array)
}

func textoMulti() {
	texto := `Me llamo Agustin
	tengo 23 años
	Juego futbol`
	fmt.Println(texto)
}

func concatenarStrings() {
	prueba := "Me llamo Agustin \n"
	prueba += "tengo 30 años"
	fmt.Println(prueba)

}

func convertir() {
	prueba := "Soy el numero " + strconv.Itoa(23)
	fmt.Println(prueba)

	prueba2 := "Soy el numero " + fmt.Sprint(23)
	fmt.Println(prueba2)
}

func metodos() {
	prueba := "Hola soy Agustin"

	if strings.Contains(prueba, "soy") {
		fmt.Println("Contiene la palabra soy")
	} else {
		fmt.Println("No contiene la palabra soy")
	}

	pruebaMay := strings.ToUpper(prueba)
	fmt.Println(pruebaMay)

	pruebaMin := strings.ToLower(prueba)
	fmt.Println(pruebaMin)

}

func busqueda() {
	prueba := "Hola soy Agustin"
	pos := strings.Index(prueba, "Agus")
	fmt.Println(pos)
}

func borrarEspacios() {
	prueba := " Hola soy Agustin "
	fmt.Println(strings.TrimSpace(prueba))

	array := strings.Split(prueba, " ")
	fmt.Println(array[1])
}
