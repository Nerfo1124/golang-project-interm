package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	// Capturando valor y error
	myValue, err := strconv.ParseInt("NaN", 0, 64)

	// Validando errores.
	if err != nil {
		fmt.Println("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// Mapa clave valor.
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	// Slice de enteros.
	s := []int{1, 2, 3}
	fmt.Println("Imprimiendo slice de enteros")
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	s = append(s, 16)
	fmt.Println("Imprimiendo slice despues de agregar valor")
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
}
