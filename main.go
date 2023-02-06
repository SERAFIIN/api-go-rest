package main

import (
	"fmt"

	"github.com/SERAFIIN/go-rest-api/database"
	"github.com/SERAFIIN/go-rest-api/models"
	"github.com/SERAFIIN/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "História 1"},
		{Id: 2, Nome: "Nome 2", Historia: "História 2"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando Servidor Rest com Go")
	routes.HandleRequest()
}
