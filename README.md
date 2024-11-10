# Projeto de Sistema de Crédito
Este repositório contém a implementação de um sistema simples que calcula o score de crédito de um usuário com base em dois fatores principais: bens (imóveis, veículos) e dívidas. A API oferece funcionalidades tanto para usuários comuns quanto para administradores, que podem consultar o score de crédito de outros usuários.

# Visão Geral do Projeto
Este sistema possui algumas funcionalidades principais:

Cálculo do Score de Crédito: A API calcula o score de crédito de um usuário, baseado na quantidade de bens que ele possui e nas dívidas que ele tem. Quanto mais bens (imóveis, veículos), maior o score. Quanto mais dívidas e de maior valor, menor o score.

Autenticação e Autorização: A API exige um token JWT para validar as requisições. Apenas usuários autenticados podem acessar as informações de score de crédito.

API para Administradores: Administradores podem consultar o score de crédito de qualquer usuário, enquanto usuários comuns só podem ver o seu próprio score.

# Estrutura do Projeto
Este projeto segue uma arquitetura simples e limpa, dividida nas seguintes pastas:

```
/root
│
├── /cmd
│   └── /api
│       └── main.go           # Ponto de entrada da aplicação
│
├── /database                 # Acesso ao banco de dados (pode conter scripts ou integrações)
│   └── database.go           # Configuração de conexão com o banco de dados
│
├── /domain                   # Definições do domínio, onde estão os modelos e lógica central
│   └── usuario.go            # Modelo de usuário
│   └── score.go              # Modelo relacionado ao cálculo de score
│
├── /handlers                 # Lógica de resposta das requisições HTTP
│   └── score.go              # Chamada as requisições HTTP
│   └── divida.go             # Chamada as requisições HTTP
│
├── /middleware               # Middleware de autenticação, autorização, etc
│   └── autenticacao.go       # Validações de autenticação
│
├── /model                    # Representação dos dados e entidades do sistema
│   └── bem.go                # Entidade de Bem no modelo
│   └── divida.go             # Entidade de Dívida no modelo
│
├── /router                   # Configuração das rotas da aplicação
│   └── router.go             # Definição das rotas
│
├── /usecases                 # Casos de uso da aplicação
│   └── bem.go                # Estruturas e regras de negocio para busca de bens, exemplo !
│
├── /utils                    # Funções auxiliares
│   └── utils.go              # Funcções auxiliares
│
├── .env                      # Variáveis de ambiente, como JWT_TOKEN
└── go.mod                    # Dependências do Go

```

# Tecnologias Utilizadas
Go (Golang): Linguagem principal usada para a construção do sistema.
JWT (JSON Web Token): Para autenticação e controle de acesso.
Banco de Dados JSON (Em memória): Usado uma estrutura simples em memória para armazenar dados. A escolha do banco JSON visa a facilidade de manipulação e a simplicidade para este projeto, mas em uma aplicação real seria mais apropriado um banco de dados mais robusto, como PostgreSQL ou MySQL.
HTTProuter: Para gerenciamento de rotas da API, uma solução rápida e simples de usar em Go.

# Como Rodar o Projeto
- Clone o repositório para o seu ambiente local
- Instale as dependências (go mod tidy)
- Configuração das Variáveis de Ambiente (JWT_TOKEN=suachavesecreta)
- Rodando a Aplicação (go run ./cmd/api/main.go)

# Motivos para o Banco JSON e Soluções Simples
Optado por um banco em memória e uma estrutura JSON devido à simplicidade do projeto e para facilitar o desenvolvimento rápido. Com o uso do JSON:
- Facilidade de manipulação: O armazenamento e acesso aos dados são rápidos e diretos.
- Simplicidade: Não há necessidade de configurar um banco de dados SQL ou NoSQL, o que acelera o desenvolvimento.
- Futuro escalável: Em uma aplicação real, o banco em memória pode ser facilmente substituído por um banco de dados mais robusto, caso a necessidade de escalabilidade surja.

# Testando a API
Para testar a API, você pode usar ferramentas como Postman ou cURL.

```
Para realizar o login e gerar o token JWT necessário para outras rotinas:
Método: GET
URL: http://localhost:8080/login
Body: '{
  "email":"teste@teste.com",
  "senha": "senha123"
}'
```

```
Para consultar o score do usuário autenticado (seu próprio score):
Método: GET
URL: http://localhost:8080/score
Header: Authorization: Bearer seu_token_jwt
```
```
Para consultar o score de outro usuário (somente administradores, com permissão do tipo 1):
Método: GET
URL: http://localhost:8080/score/admin/:id_usuario
Header: Authorization: Bearer seu_token_jwt
```