package main

import "fmt"

func MyFunc[T string | int](t T) {
	fmt.Println("t =", t)
}

/*
func main() {
	MyFunc("It works!")
}
*/
