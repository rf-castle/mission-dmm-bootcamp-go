package timelines

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/url"
	"strconv"
	"yatter-backend-go/app/app"
	"yatter-backend-go/app/domain/object"
)

// Implementation of handler
type handler struct {
	app *app.App
}

// Create Handler for `/v1/timelines/`
func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()
	h := &handler{app: app}
	r.Get("/public", h.public)
	return r
}

func parseInt64(values url.Values, key string, d int64) (int64, error) {
	c := values.Get(key)
	if c == "" {
		return d, nil
	}
	x, err := strconv.Atoi(c)
	if err != nil {
		return 0, err
	}
	return int64(x), nil
}

func parseQuery(r *http.Request) (*object.TimeLineFilter, error) {
	entity := new(object.TimeLineFilter)
	query := r.URL.Query()
	var err error
	entity.Limit, err = parseInt64(query, "limit", 40)
	if err != nil {
		return nil, err
	}
	if entity.Limit > 80 {
		entity.Limit = 80
	}
	entity.MaxId, err = parseInt64(query, "max_id", -1)
	if err != nil {
		return nil, err
	}
	entity.SinceId, err = parseInt64(query, "since_id", -1)
	if err != nil {
		return nil, err
	}
	var onlyMedia int64
	onlyMedia, err = parseInt64(query, "only_media", 0)
	if err != nil {
		return nil, err
	}
	entity.OnlyMedia = onlyMedia != 0
	return entity, nil
}
