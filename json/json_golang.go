package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	/*dir := Direccion{
		Calle:  "Belgrano",
		Numero: 120,
	}

	per := Persona{
		Nombre:   "Agustin",
		Apellido: "Alvarez",
		Edad:     23,
		Direc:    dir,
	}

	dataJson, err := json.Marshal(per)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dataJson))

	DatosMascota()
	DecodeJson()*/
}

type Direccion struct {
	Calle  string
	Numero int
}

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
	Direc    Direccion
}

type Mascota struct {
	Nombre  string   `json:"nombre"`
	Edad    int      `json:"edad,omitempty"`
	Especie string   `json:"especie,omitempty"`
	Vacunas []string `json:"vacunas"`
}

func DatosMascota() {
	m := Mascota{
		Nombre:  "Miranda",
		Edad:    4,
		Especie: "Bulldog Frances",
		Vacunas: []string{"Rabia", "Parbovirus"},
	}
	dataJson, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dataJson))
}

func DecodeJson() {
	datosJson := `{"Nombre":"Agustin","Apellido":"Alvarez","Edad":23,"Direc":{"Calle":"Belgrano","Numero":120}}`
	per := Persona{}
	err := json.Unmarshal([]byte(datosJson), &per)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", per)
}
