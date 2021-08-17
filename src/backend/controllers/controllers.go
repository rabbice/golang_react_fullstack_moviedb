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

func NewHandler() (HandlerInterface, error) {
	return DBHandler("mysql", "root:root@/movieapi")
}

func DBHandler(dbtype, conn string) (HandlerInterface, error) {
	db, err := dblayer.InitDB(dbtype, conn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		DB: db,
	}, nil
}

func HandlerWithDB(DB dblayer.DBLayer) HandlerInterface {
	return &Handler{DB: DB}
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
	id, err := strconv.Atoi(c.Request.URL.Query().Get(":id"))
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
	id, err := strconv.Atoi(c.Request.URL.Query().Get(":id"))
	if err != nil || id < 1 {
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
