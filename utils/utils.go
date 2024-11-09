package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrevalario/projeto-estudos-score/domain"
)

type Response struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type JsonResponse struct {
	Meta   interface{}       `json:"meta"`
	Data   interface{}       `json:"data"`
	Errors []domain.ApiError `json:"errors"`
}

type MetaResponse struct {
	Count int `json:"count"`
}

// SendJSONResponse envia uma resposta JSON padronizada
func SendJSONResponse(w http.ResponseWriter, data any, statusCode int) {
	// Define o tipo de conteúdo como JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Cria a resposta com a mensagem fornecida
	response := Response{Data: data}

	// Serializa o objeto Response para JSON e envia
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Erro ao enviar a resposta", http.StatusInternalServerError)
	}
}

// SendErrorResponse envia uma resposta de erro para o cliente
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	http.Error(w, fmt.Sprintf("Erro: %s", message), statusCode)
}

// DecodeRequestBody lê e decodifica o corpo da requisição para a estrutura fornecida
func DecodeRequestBody(w http.ResponseWriter, r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(target); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao ler os dados: %v", err), http.StatusBadRequest)
		return err
	}
	return nil
}

// ValidationJsonResponse retorna erros de validação
func ValidationJsonResponse(ctx context.Context, w http.ResponseWriter, errors []domain.ApiError) {
	statusCode := http.StatusBadRequest
	if len(errors) > 0 {
		statusCode = errors[0].Status
	}

	errResponse := NewJsonResponse(nil, errors, MetaResponse{
		Count: len(errors),
	})

	SendJSONResponse(w, errResponse, statusCode)
}

// NewJsonResponse formata a estrutra de retorno de slice de erros
func NewJsonResponse(data interface{}, errors []domain.ApiError, meta interface{}) JsonResponse {
	if errors == nil {
		errors = []domain.ApiError{}
	}

	if data == nil {
		data = []interface{}{}
	}

	response := JsonResponse{
		Data:   data,
		Meta:   meta,
		Errors: errors,
	}

	return response
}

func ErrorResponseJson(ctx context.Context, w http.ResponseWriter, err error) {
	errResponse, status := formatErrResponse(err)

	SendJSONResponse(w, errResponse, status)
}

func formatErrResponse(err error) (res JsonResponse, status int) {
	var errors []domain.ApiError
	apiError, ok := err.(*domain.ApiError)

	if ok {
		errors = []domain.ApiError{*apiError}
		status = apiError.Status
	} else {
		errors = []domain.ApiError{{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		}}
		status = http.StatusInternalServerError
	}

	return NewJsonResponse(nil, errors, MetaResponse{
		Count: len(errors),
	}), status
}
