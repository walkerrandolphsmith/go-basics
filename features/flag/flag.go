package flag

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/walkerrandolphsmith/go-playground/internal/config"
)

type Config struct {
	*config.Config
}

type Flag struct {
	Slug string `json:"slug"`
	Title string `json:"title"`
}

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{flagID}", GetByID(configuration))
	router.Delete("/{flagID}", DeleteByID(configuration))
	router.Post("/", Create(configuration))
	router.Get("/", GetAll(configuration))
	return router
}

func GetByID (configuration *config.Config) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		flagID := chi.URLParam(r, "flagID")
		flags := Flag{
			Slug: flagID,
			Title: "Fake Feature Flag",
		}
		render.JSON(w, r, flags)
	}
}

func DeleteByID (configuration *config.Config) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		// flagId = chi.URLParam(r, "flagID")
		response := make(map[string]string)
		response["message"] = "Deleted Successfully"
		render.JSON(w, r, response)
	}
}

func Create (configuration *config.Config) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Created Successfully"
		render.JSON(w, r, response)
	}
}

func GetAll (configuration *config.Config) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		flags := []Flag{
			{
				Slug: "slug",
				Title: "Title",
			},
		}
		render.JSON(w, r, flags)
	}
}