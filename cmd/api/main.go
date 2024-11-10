package main

import (
	"log"

	"github.com/andrevalario/projeto-estudos-score/router"
	"github.com/joho/godotenv"
)

func main() {
	startApp()
}

func startApp() {
	// Inicia servidor
	loadEnv()
	router.LoadServer()

}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	log.Println(".env carregado com sucesso")
}
