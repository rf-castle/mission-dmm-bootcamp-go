package timelines

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/handler/httperror"
)

func (h *handler) public(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	filter, err := parseQuery(r)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	repoStatus := h.app.Dao.Status()
	statuses, err := repoStatus.GetPublic(ctx, filter)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	if err = json.NewEncoder(w).Encode(statuses); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
