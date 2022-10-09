package main

import (
	"fmt"
	"sort"
)

func mapa() {
	//crear un mapa
	var meses = make(map[string]int)
	meses["Enero"] = 1
	meses["Febrero"] = 2
	meses["Marzo"] = 3
	meses["Abril"] = 4
	meses["Mayo"] = 5
	meses["Junio"] = 6

	fmt.Println(meses)

	//eliminar un elemento del mapa
	delete(meses, "Marzo")

	//recorrer un mapa
	for key, value := range meses {
		fmt.Printf("El mes %v es el numero %v \n", key, value)
	}

	//ordenar un mapa
	claves := []string{}
	for clave := range meses {
		claves = append(claves, clave)
	}

	sort.Strings(claves)

	for _, clave := range claves {
		fmt.Println(clave, meses[clave])
	}
}
