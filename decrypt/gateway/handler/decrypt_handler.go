package handler

import (
	"net/http"

	"github.com/prongbang/mock-server/decrypt/domain"
)

// DecryptHandler is the interface
type DecryptHandler interface {
	Filter(w http.ResponseWriter, r *http.Request)
}

type decryptHandler struct {
	GetData domain.GetDataUseCase
}

// NewDecryptHandler is function create instance
func NewDecryptHandler(getData domain.GetDataUseCase) DecryptHandler {
	return &decryptHandler{
		GetData: getData,
	}
}

func (h *decryptHandler) Filter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
