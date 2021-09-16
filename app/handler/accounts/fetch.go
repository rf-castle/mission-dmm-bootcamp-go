package accounts

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"yatter-backend-go/app/handler/httperror"
)

// Handle request for `GET /v1/accounts/{username}`
func (h *handler) Fetch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := chi.URLParam(r, "username")
	repoAccount := h.app.Dao.Account()
	account, err := repoAccount.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
	}
	// Todo: When account is nil, then not found user
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
