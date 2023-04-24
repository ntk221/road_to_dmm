package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Username string
	Password string
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	account := new(object.Account)
	account.Username = req.Username
	if err := account.SetPassword(req.Password); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	ret := h.app.Dao.Account() // domain/repository の取得

	if user, err := ret.FindByUsername(ctx, account.Username); err != nil {
		// httperror.Conflict(w, err)
		httperror.BadRequest(w, err)
		return
	} else if user != nil {
		// httperror.Conflict(w, err)
		httperror.BadRequest(w, err)
		return
	} else {
		if err := ret.Create(ctx, account); err != nil {
			httperror.InternalServerError(w, err)
			return
		}
	}

	// panic("Must Implement Account Registration")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}

// Handle request for `GET /v1/accounts/{username}`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")

	account, err := h.app.Dao.Account().FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// panic("Must Implement Account Get")

}
