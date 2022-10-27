package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // enviar sum al canal
}

func main() {
	/*s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // recepcion del canal

	fmt.Println(x, y, x+y)*/
	/*nombres := make(chan string, 3)
	nombres <- "juan"
	nombres <- "pedro"
	nombres <- "alberto"
	CanalesReadWrite(nombres)
	EjemploBuffer2()*/

	//creamos los 3 canales
	canal1 := make(chan string, 1)
	canal2 := make(chan string, 1)
	canal3 := make(chan string, 1)
	//llamamos a las go-rutinas para que se ejecuten
	//al mismo tiempo las 3 funciones
	go Random1(canal1)
	go Random2(canal2)
	go Random3(canal3)
	//sentencia select para que tome al proceso
	//que termino primero
	select {
	case ganador := <-canal1:
		fmt.Println(ganador)
	case ganador := <-canal2:
		fmt.Println(ganador)
	case ganador := <-canal3:
		fmt.Println(ganador)
	case <-time.After(time.Second * 25):
		fmt.Println("Erorr timeout...")
	}
}
