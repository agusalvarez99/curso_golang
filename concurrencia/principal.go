package main

import (
	"fmt"
	"time"
)

func main() {
	//go rutina que retorna algun valor
	go func() {
		res := procesoLargo()
		fmt.Println(res)
	}()
	//go rutina que no devuelve ningun valor
	go otroProcesoLargo()
	time.Sleep(time.Second * 10)
	fmt.Println("Fin del programa!")
}

func procesoLargo() string {
	fmt.Println("comenzo el proceso 1")
	//va a tardar 10 segundos en continuar la ejecucion
	time.Sleep(time.Second * 10)
	return "termino el proceso 1"
}

func otroProcesoLargo() {
	fmt.Println("comenzo el proceso 2")
	time.Sleep(time.Second * 10)
	fmt.Println("termino el proceso 2")

}
