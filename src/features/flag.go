package flag

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Flag struct {
	Slug string `json:"slug"`
	Title string `json:"title"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{flagID}", GetByID)
	router.Delete("/{flagID}", DeleteByID)
	router.Post("/", Create)
	router.Get("/", GetAll)
	return router
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	flagID := chi.URLParam(r, "flagID")
	flags := Flag{
		Slug: flagID,
		Title: "Fake Feature Flag",
	}
	render.JSON(w, r, flags)
}

func DeleteByID(w http.ResponseWriter, r *http.Request) {
	// flagId = chi.URLParam(r, "flagID")
	response := make(map[string]string)
	response["message"] = "Deleted Successfully"
	render.JSON(w, r, response)
}

func Create(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Successfully"
	render.JSON(w, r, response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	flags := []Flag{
		{
			Slug: "slug",
			Title: "Title",
		},
	}
	render.JSON(w, r, flags)
}