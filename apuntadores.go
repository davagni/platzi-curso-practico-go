package main

import "fmt"

func main() {
	x := 25

	fmt.Println(x)
	//Pasa la refencia de 'x'
	fmt.Println(&x)

	y := &x
	fmt.Println(y)
	//Trae el valor de la refencia
	fmt.Println(*y)
}

func cambiarValor(a int) {
	fmt.Println(&a)
	a = 36
}
