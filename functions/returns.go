package main

import "fmt"

// Esta funcion no me permite sumar mas de dos numeros
func sum(a, b int) int {
	return a + b
}

// Funcion variadica
func suma(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// Retornos con nombres
func getValues(x int) (int, int, int) {
	return 2 * x, 3 * x, 4 * x
}

func getValues2(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1, 2))

	fmt.Println(suma(1, 2))
	fmt.Println(suma(1, 2, 3))
	fmt.Println(suma(1, 2, 3, 4, 5))

	printNames("Alice", "Bob", "Nerfo", "Juan")

	fmt.Println(getValues(2))
	fmt.Println(getValues2(3))
}
