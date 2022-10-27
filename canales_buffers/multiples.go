package main

import (
	"math/rand"
	"time"
)

func Random1(canal chan string) {
	time.Sleep(time.Second * time.Duration(rand.Intn(100)))
	canal <- "Apuesta 1"
}

func Random2(canal chan string) {
	time.Sleep(time.Second * time.Duration(rand.Intn(100)))
	canal <- "Apuesta 2"
}

func Random3(canal chan string) {
	time.Sleep(time.Second * time.Duration(rand.Intn(100)))
	canal <- "Apuesta 3"
}
