package router

import (
	"fmt"
	"net/http"

	"github.com/andrevalario/projeto-estudos-score/handlers"
	mdlmiddleware "github.com/andrevalario/projeto-estudos-score/middleware"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/julienschmidt/httprouter"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Função responsável por iniciar o servidor
func LoadServer() {
	router := LoadRouter()
	apiPort := ":8081"
	fmt.Printf("Iniciando servidor na porta %s\n", apiPort)

	// Usa o router carregado pelo LoadRouter
	err := http.ListenAndServe(apiPort, router)
	if err != nil {
		fmt.Printf("Erro ao iniciar servidor: %v\n", err)
	}
}

// Função que carrega as rotas
func LoadRouter() http.Handler {
	// Configura o router e as rotas
	router := httptrace.New(
		httptrace.WithServiceName("projeto-estudo-score"),
		httptrace.WithSpanOptions(
			tracer.Tag(ext.SamplingPriority, ext.PriorityUserKeep),
		),
	)

	// Adiciona a rota de verificação de vida
	router.GET("/alive", handlers.Alive)

	// Chama as funções de configuração das rotas
	loadRouterAutenticacao(router)
	loadRouterUsuario(router)
	loadRouterDividas(router)
	loadRouterBens(router)
	loadRouterScore(router)

	return router
}

// Função que configura as rotas relacionadas á autenticação
func loadRouterAutenticacao(router *httptrace.Router) {
	router.POST("/login", handlers.LoginUsuario)
}

// Função que configura as rotas relacionadas ao usuário
func loadRouterUsuario(router *httptrace.Router) {
	router.POST("/usuario", handlers.CriarUsuario)
	router.GET("/usuarios/:id", handlers.BuscarUsuarioPorID)
	router.PUT("/usuarios/:id", handlers.AtualizarUsuario)
	router.DELETE("/usuarios/:id", handlers.DeletarUsuario)
}

// Função que configura as rotas relacionadas as dividas de um usuário
func loadRouterDividas(router *httptrace.Router) {
	router.POST("/dividas", mdlmiddleware.ValidarToken(mdlmiddleware.ValidarAcessoDivida(handlers.CriarDivida)))
	router.GET("/dividas/:id", mdlmiddleware.ValidarToken(mdlmiddleware.ValidarAcessoDivida(handlers.BuscarDivida)))
	router.PUT("/dividas/:id", mdlmiddleware.ValidarToken(mdlmiddleware.ValidarAcessoDivida(handlers.AtualizarDivida)))
	router.DELETE("/dividas/:id", mdlmiddleware.ValidarToken(mdlmiddleware.ValidarAcessoDivida(handlers.DeletarDivida)))
}

// Função que configura as rotas relacionadas aos bens vinculados a um usuário
func loadRouterBens(router *httptrace.Router) {
	router.POST("/bens", mdlmiddleware.ValidarToken(mdlmiddleware.ValidarAcessoBens(handlers.CriarBem)))
	router.GET("/bens/:id", mdlmiddleware.ValidarToken(mdlmiddleware.ValidarAcessoBens(handlers.BuscarBemPorID)))
	router.PUT("/bens/:id", mdlmiddleware.ValidarToken(mdlmiddleware.ValidarAcessoBens(handlers.AtualizarBem)))
	router.DELETE("/bens/:id", mdlmiddleware.ValidarToken(mdlmiddleware.ValidarAcessoBens(handlers.DeletarBem)))
}

func loadRouterScore(router *httptrace.Router) {
	router.GET("/score/admin/:id_usuario", mdlmiddleware.ValidarToken(mdlmiddleware.ValidarAcessoAdmin(handlers.CalcularScoreAdmin)))
	router.GET("/score", mdlmiddleware.ValidarToken(handlers.CalcularScoreUsuario))
}
