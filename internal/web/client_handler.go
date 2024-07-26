package web

import (
	"encoding/json"
	"fmt"
	create_client "github.com/thgbianeck/bnck-ms-wallet/internal/usecase/createClient"
	"net/http"
)

type WebClientHandler struct {
	CreateClientUseCase create_client.CreateClientUseCase
}

func NewWebClientHandler(createClientUseCase create_client.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{CreateClientUseCase: createClientUseCase}
}

func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto create_client.CreateClientInputDTO
	fmt.Println(dto)
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	output, err := h.CreateClientUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
