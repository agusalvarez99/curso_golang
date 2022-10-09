package main

import "fmt"

type Persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	/*persona_1 := Persona{
		nombre:   "Agustin",
		apellido: "Alvarez",
		edad:     23,
	}
	persona_2 := Persona{
		nombre:   "Melisa",
		apellido: "Alvarez",
		edad:     24,
	}

	persona_3 := Persona{
		nombre:   "Agustin",
		apellido: "Alvarez",
		edad:     23,
	}

	if persona_1 != persona_2 {
		fmt.Println("Las personas 1 y 2 son difentes")
	}

	if persona_1 == persona_3 {
		fmt.Println("Las personas 1 y 3 son iguales")
	}*/

	conSQL := ConexionSQL{Puerto: 3306}
	conPracle := ConexionOracle{Puerto: 1500}

	DatosConexion(&conSQL)
	DatosConexion(&conPracle)
}

type Conexion interface {
	Conectar() string
	Desconectar() string
}

type ConexionSQL struct {
	Puerto int
}

func (con *ConexionSQL) Conectar() string {
	return "Conexion SQL exitosa"
}

func (con *ConexionSQL) Desconectar() string {
	return "Conexion SQL cerrada"
}

type ConexionOracle struct {
	Puerto int
}

func (con *ConexionOracle) Conectar() string {
	return "Conexion Oracle exitosa"
}

func (con *ConexionOracle) Desconectar() string {
	return "Conexion Oracle cerrada"
}

func DatosConexion(c Conexion) {
	fmt.Println(c.Conectar())
	fmt.Println(c.Desconectar())
}
