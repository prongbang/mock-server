package route

import (
	"log"
	"net/http"

	"github.com/prongbang/mock-server/decrypt/gateway/handler"
)

// DecryptRoute is the interface
type DecryptRoute interface {
	Initial()
}

type decryptRoute struct {
	Handler handler.DecryptHandler
}

// NewDecryptRoute is function for create instance
func NewDecryptRoute(handler handler.DecryptHandler) DecryptRoute {
	return &decryptRoute{
		Handler: handler,
	}
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s -> %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func (route *decryptRoute) Initial() {

	http.Handle("/", withLogging(http.HandlerFunc(route.Handler.Filter)))
}
