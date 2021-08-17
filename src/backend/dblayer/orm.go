package dblayer

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rabbice/movieapi/src/backend/models"
)

type DBORM struct {
	*gorm.DB
}

func InitDB(dbname, conn string) (*DBORM, error) {
	db, err := gorm.Open(dbname, conn+"?parseTime=true")
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllMovies() (movies []models.Movie, err error) {
	return movies, db.Find(&movies).Error
}

func (db *DBORM) GetMovieByID(id int) (movies models.Movie, err error) {
	return movies, db.Where("id = ?", id).First(&movies).Error
}

func (db *DBORM) AddMovie(movies models.Movie) (models.Movie, error) {
	return movies, db.Create(&movies).Error
}

func (db *DBORM) DeleteMovieByID(id int) (movies models.Movie, err error) {
	return movies, db.Where("id = ?", id).Delete(&movies).Error
}

func (db *DBORM) UpdateMovieByID(id int) (movies models.Movie, err error) {
	return movies, db.Model(&models.Movie{}).Where("id = ?", id).First(&movies).Update(&movies.Title).Error
}