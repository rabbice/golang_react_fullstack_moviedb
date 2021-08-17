package dblayer

import (
	"github.com/rabbice/movieapi/src/backend/models"
)

type DBLayer interface {
	GetAllMovies() ([]models.Movie, error)
	GetMovieByID(int) (models.Movie, error)
	AddMovie(models.Movie) (models.Movie, error)
	DeleteMovieByID(int) (models.Movie, error)
}
