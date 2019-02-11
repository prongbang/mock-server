package main

import (
	"net/http"

	"github.com/prongbang/mock-server/decrypt/di"
)

func main() {

	route := di.ProvideDecryptRoute()
	route.Initial()

	http.ListenAndServe(":3000", nil)
}
