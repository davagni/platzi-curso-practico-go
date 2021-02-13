package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	canal := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	tiempoPasado := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion: %s\n", tiempoPasado)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "No está disponible =(")
		canal <- servidor + " no se encuentra disponible"
	} else {
		fmt.Println(servidor, "Funcionando normalmente =)")
		canal <- servidor + " funcionando normalmente"
	}
}
