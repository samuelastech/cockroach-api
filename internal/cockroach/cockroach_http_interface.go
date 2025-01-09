package cockroach

import "net/http"

type CockroachHandler interface {
	CreateCockroach(w http.ResponseWriter, r *http.Request) error
}
