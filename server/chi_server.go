package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samulastech/cockroach/config"
	"github.com/samulastech/cockroach/internal/cockroach"
	"log"
	"net/http"
)

type ChiServer struct {
	app  *chi.Mux
	conf *config.Config
}

func NewChiServer(conf *config.Config) *ChiServer {
	return &ChiServer{
		app:  chi.NewRouter(),
		conf: conf,
	}
}

func (server *ChiServer) Start() error {
	server.app.Use(middleware.Logger)
	server.app.Use(middleware.Recoverer)

	server.app.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	server.initCockroachHttpHandler()

	return http.ListenAndServe(fmt.Sprintf(":%s", server.conf.Server.Port), server.app)
}

func (server *ChiServer) initCockroachHttpHandler() {
	repository := cockroach.NewCockroachRepositoryInMemo()
	messenger := cockroach.NewCockroachMessaging()
	usecaseCreate := cockroach.NewCockroachUsecaseCreate(repository, messenger)
	handler := cockroach.NewCockroachHTTPHandler(usecaseCreate)

	server.app.Post("/cockroach", handler.CreateCockroach)
	log.Println("[message: cockroach routes initialized]")
}
