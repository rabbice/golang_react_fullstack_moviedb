package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/movieapi/src/backend/controllers"
)

func RunAPI(address string) error {
	m, err := controllers.NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, m)
}

func RunAPIWithHandler(address string, m controllers.HandlerInterface) error {
	r := gin.Default()
	r.GET("/movies", m.GetMovies)
	r.GET("/movie/:id", m.ShowMovie)
	r.POST("/movie/create", m.AddMovie)
	r.DELETE("movie/:id/delete", m.DeleteMovie)
	return r.Run(address)
}
