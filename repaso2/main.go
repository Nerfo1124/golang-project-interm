package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	go doSomething(c)
	<-c

	g := 25
	fmt.Println(g) // imprime el valor entero 25
	h := &g
	fmt.Println(h) // imprimer la direccion de memoria.
	i := *h
	fmt.Println(i) // Imprime el valor por de g
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("done")
	c <- 1
}
