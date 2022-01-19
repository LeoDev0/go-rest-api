package main

import (
	"fmt"

	"github.com/leodev0/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}