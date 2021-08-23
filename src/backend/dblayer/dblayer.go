package dblayer

import (
	"errors"

	"github.com/rabbice/movieapi/src/backend/models"
)

var ErrINVALIDPASSWORD = errors.New("invalid password")

type DBLayer interface {
	GetAllMovies() ([]models.Movie, error)
	GetMovieByID(int) (models.Movie, error)
	AddMovie(models.Movie) (models.Movie, error)
	DeleteMovieByID(int) (models.Movie, error)
	UpdateMovieByID(int) (models.Movie, error)
	AddUser(models.User) (models.User, error)
	SignInUser(email, password string) (models.User, error)
	SignOutUserById(int) error
	AuthUser(models.User) (models.User, error)
}
