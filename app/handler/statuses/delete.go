package statuses

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
	"yatter-backend-go/app/handler/request"
)

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := request.IDOf(r)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	ctx := r.Context()
	repoAccount := h.app.Dao.Account()
	repoStatus := h.app.Dao.Status(repoAccount)
	account := auth.AccountOf(r)
	status, err := repoStatus.FindById(ctx, id)
	if err != nil {
		// Todo: Not found error
		httperror.InternalServerError(w, err)
		return
	}
	if account.ID != status.Account.ID {
		// Not Authorized
		httperror.BadRequest(w, errors.New("You cannot delete this status"))
		return
	}
	if err = repoStatus.Delete(ctx, status.ID); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// これでエラーが出るとは思えんが・・・
	if err = json.NewEncoder(w).Encode(struct{}{}); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
