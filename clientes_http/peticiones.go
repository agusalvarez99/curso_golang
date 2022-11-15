package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//get()
	//post()
	//put()
	delete()
}

func get() {
	clienteHttp := &http.Client{}
	url := "https://httpbin.org/get"
	peticion, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	peticion.Header.Add("X-Prueba", "Ejemplo")
	respuesta, err := clienteHttp.Do(peticion)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	defer respuesta.Body.Close()

	respuestaBytes, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	respuestaString := string(respuestaBytes)
	log.Printf("Response code: %d", respuesta.StatusCode)
	log.Printf("Response body: '%s'", respuestaString)
}

func post() {
	clienteHttp := &http.Client{}
	url := "https://httpbin.org/post"

	type Animal struct {
		Nombre  string
		Especie string
	}
	ani := Animal{
		Nombre:  "Miru",
		Especie: "Perro",
	}

	aniJson, err := json.Marshal(ani)
	if err != nil {
		log.Fatalf("Error con el JSON %v", err)
	}

	peticion, err := http.NewRequest("POST", url, bytes.NewBuffer(aniJson))
	if err != nil {
		log.Fatalf("Error en peticion %v", err)
	}

	peticion.Header.Add("Content-Type", "application/json")
	respuesta, err := clienteHttp.Do(peticion)
	if err != nil {
		log.Fatalf("Error en peticion %v", err)
	}
	defer respuesta.Body.Close()
	body, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	res := string(body)
	log.Printf("Response code: %d", respuesta.StatusCode)
	log.Printf("Headers: '%q'", respuesta.Header)
	contentType := respuesta.Header.Get("Content-Type")
	log.Printf("Content-type: '%s'", contentType)
	log.Printf("Response body: '%s'", res)
}

func put() {
	clienteHttp := &http.Client{}
	url := "https://httpbin.org/put"

	type Animal struct {
		Nombre  string
		Especie string
	}
	ani := Animal{
		Nombre:  "Miru",
		Especie: "Perro",
	}

	aniJson, err := json.Marshal(ani)
	if err != nil {
		log.Fatalf("Error con el JSON %v", err)
	}

	peticion, err := http.NewRequest("PUT", url, bytes.NewBuffer(aniJson))
	if err != nil {
		log.Fatalf("Error en peticion %v", err)
	}

	peticion.Header.Add("Content-Type", "application/json")
	respuesta, err := clienteHttp.Do(peticion)
	if err != nil {
		log.Fatalf("Error en peticion %v", err)
	}
	defer respuesta.Body.Close()
	body, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	res := string(body)
	log.Printf("Response code: %d", respuesta.StatusCode)
	log.Printf("Headers: '%q'", respuesta.Header)
	contentType := respuesta.Header.Get("Content-Type")
	log.Printf("Content-type: '%s'", contentType)
	log.Printf("Response body: '%s'", res)
}

func delete() {
	clienteHttp := &http.Client{}
	url := "https://httpbin.org/delete"

	type Animal struct {
		Nombre  string
		Especie string
	}
	ani := Animal{
		Nombre:  "Miru",
		Especie: "Perro",
	}

	aniJson, err := json.Marshal(ani)
	if err != nil {
		log.Fatalf("Error con el JSON %v", err)
	}
	peticion, err := http.NewRequest("DELETE", url, bytes.NewBuffer(aniJson))
	if err != nil {
		log.Fatalf("Error en peticion %v", err)
	}
	peticion.Header.Add("Content-Type", "application/json")
	respuesta, err := clienteHttp.Do(peticion)
	if err != nil {
		log.Fatalf("Error en peticion %v", err)
	}
	defer respuesta.Body.Close()

	body, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	res := string(body)
	log.Printf("Response code: %d", respuesta.StatusCode)
	log.Printf("Headers: '%q'", respuesta.Header)
	contentType := respuesta.Header.Get("Content-Type")
	log.Printf("Content-type: '%s'", contentType)
	log.Printf("Response body: '%s'", res)
}
