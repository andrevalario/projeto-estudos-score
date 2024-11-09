package main

import (
	"log"

	"github.com/andrevalario/projeto-estudos-score/router"
)

func startApp() {
	router.InitializeRouter()
}

func main() {
	startApp()
	log.Println("Servidor iniciado na porta 8080")
}
