package router

import (
	"fmt"
	"net/http"

	"github.com/andrevalario/projeto-estudos-score/handlers"
	"github.com/gorilla/mux"
)

// Função responsável por iniciar o servidor
func LoadServer() {
	// Carrega o roteador e inicia o servidor
	router := LoadRouter()
	apiPort := ":8080"
	fmt.Printf("Iniciando servidor na porta %s\n", apiPort)

	// Usa o router carregado pelo LoadRouter
	err := http.ListenAndServe(apiPort, router)
	if err != nil {
		fmt.Printf("Erro ao iniciar servidor: %v\n", err)
	}
}

// Função que carrega as rotas
func LoadRouter() *mux.Router {
	// Configura o router e as rotas
	router := mux.NewRouter()

	// Adiciona a rota de verificação de vida
	router.HandleFunc("/alive", handlers.Alive).Methods("GET")

	// Chama as funções de configuração das rotas
	loadRouterUsuario(router)
	loadRouterDividas(router)
	loadRouterBens(router)

	return router
}

// Função que configura as rotas relacionadas ao usuário
func loadRouterUsuario(router *mux.Router) {
	// Rotas para o CRUD de usuário
	router.HandleFunc("/usuario", handlers.CriarUsuario).Methods("POST")
	router.HandleFunc("/usuarios/{id}", handlers.BuscarUsuarioPorID).Methods("GET")
	router.HandleFunc("/usuarios/{id}", handlers.AtualizarUsuario).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", handlers.DeletarUsuario).Methods("DELETE")
}

// Função que configura as rotas relacionadas as dividas de um usuário
func loadRouterDividas(router *mux.Router) {
	router.HandleFunc("/dividas", handlers.CriarDivida).Methods("POST")
	router.HandleFunc("/dividas/{id}", handlers.BuscarDivida).Methods("GET")
	router.HandleFunc("/dividas/{id}", handlers.AtualizarDivida).Methods("PUT")
	router.HandleFunc("/dividas/{id}", handlers.DeletarDivida).Methods("DELETE")
}

// Função que configura as rotas relacionadas aos bens vinculados a um usuário
func loadRouterBens(router *mux.Router) {
	router.HandleFunc("/bens", handlers.CriarBem).Methods("POST")
	router.HandleFunc("/bens/{id}", handlers.BuscarBemPorID).Methods("GET")
	router.HandleFunc("/bens/{id}", handlers.AtualizarBem).Methods("PUT")
	router.HandleFunc("/bens/{id}", handlers.DeletarBem).Methods("DELETE")
}
