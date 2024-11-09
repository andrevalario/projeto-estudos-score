# projeto-estudos-score
Este projeto tem como objetivo criar uma API simples de score de crédito, com foco no aprendizado e aplicação de conceitos de backend e boas práticas no desenvolvimento de APIs utilizando Go (Golang).

## Estrutura do projeto:
```
├── cmd/
│   └── api/
│       └── main.go                # Função main que inicializa a aplicação
│
├── internal/
│   ├── handlers/                  # Lógica para manipular requisições HTTP
│   │   └── score_handler.go       # Manipulador para requisições relacionadas ao Score
│   │   └── user_handler.go        # Manipulador para requisições relacionadas ao Usuário
│   ├── usecases/                  # Lógica de negócios/funcionalidades
│   │   └── score_usecase.go       # Cálculo do score de crédito
│   │   └── user_usecase.go        # Lógica de manipulação de dados do usuário
│   ├── models/                    # Estruturas de dados, interações com o banco
│   │   └── score_model.go         # Estrutura para Score e interações com DB
│   │   └── user_model.go          # Estrutura para usuário e interações com DB
│   ├── domains/                   # Definições dos dados e entidades de domínio
│   │   └── score.go               # Definição de entidade de Score
│   │   └── user.go                # Definição de entidade de Usuário
│
├── go.mod                         # Arquivo de dependências
└── README.md                      # Documentação do projeto
```
