package di

import (
	"github.com/prongbang/mock-server/decrypt/data/repository"
	"github.com/prongbang/mock-server/decrypt/domain"
	"github.com/prongbang/mock-server/decrypt/gateway/handler"
	"github.com/prongbang/mock-server/decrypt/gateway/route"
)

// ProvideDecryptRepository is provide function
func ProvideDecryptRepository() repository.DecryptRepository {
	return repository.NewDecryptRepository()
}

// ProvideGetDataUseCase is provide function
func ProvideGetDataUseCase() domain.GetDataUseCase {
	return domain.NewGetDataUseCase(ProvideDecryptRepository())
}

// ProvideDecryptHandler is provide function
func ProvideDecryptHandler() handler.DecryptHandler {
	return handler.NewDecryptHandler(
		ProvideGetDataUseCase(),
	)
}

// ProvideDecryptRoute is provide function
func ProvideDecryptRoute() route.DecryptRoute {
	return route.NewDecryptRoute(ProvideDecryptHandler())
}
