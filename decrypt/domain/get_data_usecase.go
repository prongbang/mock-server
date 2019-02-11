package domain

import (
	"github.com/prongbang/mock-server/decrypt/data/repository"
)

// GetDataUseCase is interface
type GetDataUseCase interface {
}

type getDataUseCase struct {
	Repo repository.DecryptRepository
}

// NewGetDataUseCase is function create instance
func NewGetDataUseCase(repo repository.DecryptRepository) GetDataUseCase {
	return &getDataUseCase{
		Repo: repo,
	}
}
