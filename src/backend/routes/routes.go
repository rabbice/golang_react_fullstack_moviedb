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

func MovieService(address string, m controllers.Handlers) error {
	v1 := gin.Default()
	v1.Use(middleware.SecureHeaders())
	v1.GET("/v1/movies", m.GetMovies)
	v1.GET("/v1/movie/:id", m.ShowMovie)
	v1.POST("/v1/movie/create", m.AddMovie)
	v1.PUT("/v1/movie/:id", m.UpdateMovie)
	v1.DELETE("/v1/movie/:id", m.DeleteMovie)
	v1.POST("/v1/users/signup", m.AddUser)
	v1.POST("/v1/users/signin", m.SignIn)
	v1.POST("/v1/users/signout/:id", m.SignOut)
	v1.GET("/v1/user/:id", m.GetUser)
	return v1.Run(address)
}
