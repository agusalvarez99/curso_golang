package main

import (
	"fmt"
)

// vamos a leer un fichero y lanzar un error si no lo encuentra
func main() {
	//texto, err := ioutil.ReadFile("pruebax.txt")
	//si el error es distinto de null
	//if err != nil {
	//fmt.Println(err)
	//err2 := errors.New("Archivo no encontrado jeje")
	//miError(err2)
	//return
	//}
	//pasar a string para que muestre el texto como esta en el documento
	//fmt.Println(string(texto))

	//res, err := sumaConError(10, 24)
	//fmt.Println(res)
	//fmt.Println(err)

	//finEjecucion()

	res, err := divDecimal(25, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func miError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func sumaConError(num1 int, num2 int) (int, error) {
	if num1 > 100 || num2 > 100 {
		err := fmt.Errorf("No ingresar numeros mayores a 100")
		//retorno 0 como resultado y el error
		return 0, err
	}
	//retorno el resultado y null como error
	return num1 + num2, nil
}

func finEjecucion() {
	fmt.Println("Hasta aqui llega la ejecucion")
	panic("Paro la ejecucion")
}

func divDecimal(a float32, b float32) (float32, error) {
	if b == 0 {
		err := fmt.Errorf("No se puede dividir por 0")
		return 0, err
	}
	return a / b, nil
}
