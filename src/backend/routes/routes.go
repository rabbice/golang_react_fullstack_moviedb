package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/movieapi/src/backend/controllers"
	"github.com/rabbice/movieapi/src/backend/middleware"
)

func RunAPI(address string) error {
	m, err := controllers.Conn()
	if err != nil {
		return err
	}
	return MovieService(address, m)
}

func MovieService(address string, m controllers.MovieHandlers) error {
	r := gin.Default()
	r.Use(middleware.SecureHeaders())
	r.GET("/movies", m.GetMovies)
	r.GET("/movie/:id", m.ShowMovie)
	r.POST("/movie/create", m.AddMovie)
	r.PUT("/movie/:id", m.UpdateMovie)
	r.DELETE("movie/:id", m.DeleteMovie)
	return r.Run(address)
}
