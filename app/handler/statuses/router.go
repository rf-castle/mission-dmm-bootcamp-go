package statuses

import (
	"net/http"
	"yatter-backend-go/app/handler/auth"

	"yatter-backend-go/app/app"

	"github.com/go-chi/chi"
)

// Implementation of handler
type handler struct {
	app *app.App
}

// Create Handler for `/v1/statuses/`
func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()
	h := &handler{app: app}
	middleAuth := auth.Middleware(app)
	r.With(middleAuth).Post("/", h.Create)
	r.Get("/{id}", h.Fetch)
	return r
}
