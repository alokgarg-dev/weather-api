package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/alokgarg-dev/weather-api/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/", loadWeatherRoutes)

	return router
}

func loadWeatherRoutes(router chi.Router) {
	weatherHandler := &handler.Weather{}

	router.Get("/weather", weatherHandler.Get)
}
