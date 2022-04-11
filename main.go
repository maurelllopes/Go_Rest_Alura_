package main

import (
	"Go_Rest_Alura_/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando Servidor Rest com Go")
	routes.HandleRequest()
}
