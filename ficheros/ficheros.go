package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//leer()
	//escribir()
	//listarDir()
	//copiar()
	borrar()
}

func leer() {
	ficheroArray, err := ioutil.ReadFile("prueba.txt")
	if err != nil {
		log.Fatal(err)
	}
	contenido := string(ficheroArray)
	fmt.Println(contenido)
}

func escribir() {
	arrayBytes := []byte("Me llamo Agustin, estudio ingenieria!")
	err := ioutil.WriteFile("prueba2.txt", arrayBytes, 0664)
	if err != nil {
		log.Fatal(err)
	}
}

func listarDir() {
	listaFicheros, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range listaFicheros {
		fmt.Println("Nombre archivo:", f.Name())
		fmt.Println("Tamaño:", f.Size(), "bytes")
		fmt.Println("Es directorio?", f.IsDir())
	}

}

func copiar() {
	fichero, err := os.Open("prueba.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fichero.Close()

	copiaFichero, err := os.OpenFile("copia.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer copiaFichero.Close()

	resultado, err := io.Copy(copiaFichero, fichero)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resultado)
}

func borrar() {
	err := os.Remove("copia.txt")
	if err != nil {
		log.Fatal(err)
	}
}
