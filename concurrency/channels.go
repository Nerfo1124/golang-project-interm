package main

import "fmt"

func main() {
	//Channel sin buffer
	//c := make(chan int)

	//Channel con buffer
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

//func goFunction() {
//	<-c
//}
