package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rabbice/movieapi/src/backend/dblayer"
	"github.com/rabbice/movieapi/src/backend/models"
)

type HandlerInterface interface {
	GetMovies(c *gin.Context)
	ShowMovie(c *gin.Context)
	AddMovie(c *gin.Context)
	DeleteMovie(c *gin.Context)
}

type Handler struct {
	DB dblayer.DBLayer
}

func NewHandler() (*Handler, error) {
	return new(Handler), nil
}

func (m *Handler) GetMovies(c *gin.Context) {
	if m.DB == nil {
		return
	}
	movies, err := m.DB.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (m *Handler) ShowMovie(c *gin.Context) {
	if m.DB == nil {
		return
	}
	p := c.Param(":id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	movies, err := m.DB.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (m *Handler) AddMovie(c *gin.Context) {
	if m.DB == nil {
		return
	}
	var movie models.Movie
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	movie, err = m.DB.AddMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (m *Handler) DeleteMovie(c *gin.Context) {
	if m.DB == nil {
		return
	}
	p := c.Param(":id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	movie, err := m.DB.DeleteMovieByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}
