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

func New(configuration *config.Config) *Config {
	return &Config{configuration}
}

func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{flagID}", config.GetByID)
	router.Delete("/{flagID}", config.DeleteByID)
	router.Post("/", config.Create)
	router.Get("/", config.GetAll)
	return router
}

func (config *Config) GetByID(w http.ResponseWriter, r *http.Request) {
	flagID := chi.URLParam(r, "flagID")
	flags := Flag{
		Slug: flagID,
		Title: "Fake Feature Flag",
	}
	render.JSON(w, r, flags)
}

func (config *Config)  DeleteByID(w http.ResponseWriter, r *http.Request) {
	// flagId = chi.URLParam(r, "flagID")
	response := make(map[string]string)
	response["message"] = "Deleted Successfully"
	render.JSON(w, r, response)
}

func (config *Config)  Create(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Successfully"
	render.JSON(w, r, response)
}

func (config *Config)  GetAll(w http.ResponseWriter, r *http.Request) {
	flags := []Flag{
		{
			Slug: "slug",
			Title: "Title",
		},
	}
	render.JSON(w, r, flags)
}