package main

import (
	"errors"
	"log"
	"os"
)

func main() {
	//muestra el mensaje añadiendo fecha y hora
	log.Println("Hola mundo log!")

	err := errors.New("Error de prueba!")
	log.Println(err)

	//muestra el mensaje y detiene ejecucion
	//log.Fatal("Error fatal!")

	//escribir log en un archivo, en caso de no existir lo crea
	f, err := os.OpenFile("Prueba.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	//defer se utiliza para ejecutar algo una vez que main acabo
	defer f.Close()

	log.SetOutput(f)

	for i := 0; i <= 50; i++ {
		log.Printf("Pueba escritura en log %v", i)
	}
}
