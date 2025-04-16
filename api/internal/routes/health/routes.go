package health

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var Routes = struct {
	LBHeartbeat string
}{
	LBHeartbeat: "/__lbheartbeat__",
}

// RegisterRoutes of the health package registers health endpoints that implement Dockerflow -
// see https://github.com/mozilla-services/Dockerflow#containerized-app-requirements.
func RegisterRoutes(router *chi.Mux) {
	router.Get(Routes.LBHeartbeat, func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("OK")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
