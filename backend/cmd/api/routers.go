package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(app.enableCORS)
	// ping은 route가 하나라도 설정이 되어 있어야 동작한다.
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/", app.Home)

	// mux.Get("/refresh", app.refreshToken)
	// mux.Get("/logout", app.logout)

	mux.Get("/gmarket/items", app.AllGmarketItems)
	// mux.Get("/movies/{id}", app.GetMovie)

	// mux.Get("/genres", app.AllGenres)
	// mux.Get("/movies/genres/{id}", app.AllMoviesByGenre)

	// mux.Post("/graph", app.moviesGraphQL)

	// mux.Route("/admin", func(mux chi.Router){
	// 	mux.Use(app.authRequired)

	// 	mux.Get("/movies", app.MovieCatalog)
	// 	mux.Get("/movies/{id}", app.MovieForEdit)
	// 	mux.Put("/movies/0", app.InsertMovie)
	// 	mux.Patch("/movies/{id}", app.UpdateMovie)
	// 	mux.Delete("/movies/{id}", app.DeleteMovie)
	// })

	return mux
}
