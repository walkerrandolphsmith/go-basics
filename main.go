package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/walkerrandolphsmith/go-playground/features/flag"
	"github.com/walkerrandolphsmith/go-playground/internal/config"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/flag", flag.Routes(configuration))
	})

	return router
}

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error", err)
	}
	router := Routes(configuration)
	log.Println("Serving application at PORT :" + configuration.Constants.PORT)
	log.Fatal(http.ListenAndServe(":"+configuration.Constants.PORT, router))
}