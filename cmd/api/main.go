package main

import (
	"github.com/andrevalario/projeto-estudos-score/router"
)

func main() {
	startApp()
}

func startApp() {
	// Inicia servidor
	router.LoadServer()
}
