package status

import (
	"net/http"

	"yatter-backend-go/app/app"

	"github.com/go-chi/chi"
)

type handler struct {
	app *app.App
}

// Create new status
func NewRouter(app *app.App) http.Handler {
	h := &handler{app: app}
	r := chi.NewRouter()

	// r.Use(auth.Middleware(app))
	r.Post("/", h.Create)
	return r
}
