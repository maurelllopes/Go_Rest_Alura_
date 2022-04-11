package main

import (
	"Go_Rest_Alura_/models"
	"Go_Rest_Alura_/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}
	fmt.Println("Iniciando Servidor Rest com Go")
	routes.HandleRequest()
}
