package main

import (
	"fmt"

	"github.com/leodev0/go-rest-api/database"
	"github.com/leodev0/go-rest-api/routes"
)

func main() {
	database.ConexaoComBancoDeDados()

	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}