package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/handler/httperror"
	"yatter-backend-go/app/handler/request"
)

func (h *handler) Fetch(w http.ResponseWriter, r *http.Request) {

	id, err := request.IDOf(r)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	ctx := r.Context()
	repoAccount := h.app.Dao.Account()
	repoStatus := h.app.Dao.Status(repoAccount)
	status, err := repoStatus.FindById(ctx, id)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(status); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
