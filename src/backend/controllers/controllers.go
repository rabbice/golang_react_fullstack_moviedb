package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rabbice/movieapi/src/backend/dblayer"
	"github.com/rabbice/movieapi/src/backend/models"
)

const SecretKey = "secret"

type Handlers interface {
	GetMovies(c *gin.Context)
	ShowMovie(c *gin.Context)
	AddMovie(c *gin.Context)
	DeleteMovie(c *gin.Context)
	UpdateMovie(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
}

type Handler struct {
	DB dblayer.DBLayer
}

func Conn() (Handlers, error) {
	return DBHandler("mysql", "root:root@/movieapi")
}

func DBHandler(dbtype, conn string) (Handlers, error) {
	db, err := dblayer.InitDB(dbtype, conn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		DB: db,
	}, nil
}

func HandlerWithDB(DB dblayer.DBLayer) Handlers {
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
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

func (m *Handler) UpdateMovie(c *gin.Context) {
	if m.DB == nil {
		return
	}
	var movie models.Movie
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	movie, err = m.DB.UpdateMovieByID(id)
	if err != nil || id < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}
func (m *Handler) DeleteMovie(c *gin.Context) {
	if m.DB == nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
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

func (m *Handler) AddUser(c *gin.Context) {
	if m.DB == nil {
		return
	}
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user, err = m.DB.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (m *Handler) SignIn(c *gin.Context) {
	if m.DB == nil {
		return
	}
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = m.DB.SignInUser(user.Email, user.Password)
	if err != nil {
		//if the error is invalid password, return forbidden http error
		if err == dblayer.ErrINVALIDPASSWORD {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not login"})
		return
	}
	cookie, err := c.Cookie("jwt")
	if err != nil {
		cookie = "not set"
		c.SetCookie("jwt", token, 60*60*24, "/", "localhost", false, true)
	}
	c.JSON(http.StatusOK, cookie)
}

func (m *Handler) SignOut(c *gin.Context) {
	if m.DB == nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}
	err = m.DB.SignOutUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
