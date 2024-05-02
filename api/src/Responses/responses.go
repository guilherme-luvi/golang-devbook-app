package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Função de resposta genérica. Retorna resposta em Json para a requisição
func JSON(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.WriteHeader(statusCode)

	if err := json.NewEncoder(writer).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Retorna erro em formato Json
func Erro(writer http.ResponseWriter, statusCode int, err error) {
	JSON(writer, statusCode, struct {
		Err string `json:"erro"`
	}{
		Err: err.Error(),
	})
}