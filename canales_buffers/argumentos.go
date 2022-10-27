package main

import "fmt"

func CanalesReadWrite(c chan string) {
	fmt.Println(<-c)
	c <- "Agus"
	close(c)
	for i := range c {
		fmt.Println(i)
	}
}

func CanalesOnlyRead(c <-chan string) {
	fmt.Println(<-c)
	for i := range c {
		fmt.Println(i)
	}
}

func CanalesOnlyWrite(c chan<- string) {
	c <- "Agus"
}
