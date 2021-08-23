package dblayer

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rabbice/movieapi/src/backend/models"
	"golang.org/x/crypto/bcrypt"
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

func (db *DBORM) AddUser(user models.User) (models.User, error) {
	hashPassword(&user.Password)
	user.LoggedIn = true
	return user, db.Create(&user).Error
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("reference provided for hashing password is nil")
	}

	//convert password string to byte slice so that we can use it with the bcrypt package
	sBytes := []byte(*s)

	//Obtain hashed password via the GenerateFromPassword() method
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	//update password string with the hashed version
	*s = string(hashedBytes[:])
	return nil
}

func checkPassword(existingHash, incomingPass string) bool {
	//this method will return an error if the hash does not match the provided password string
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass)) == nil
}

func (db *DBORM) SignInUser(email, password string) (user models.User, err error) {
	// obtain a *gorm.DB object representing our customer's row
	result := db.Table("Users").Where(&models.User{Email: email})
	err = result.First(&user).Error
	if err != nil {
		return user, err
	}

	if !checkPassword(user.Password, password) {
		return user, ErrINVALIDPASSWORD
	}
	// update the loggedin field
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return user, err
	}
	// return the new customer row
	return user, result.Find(&user).Error
}

func (db *DBORM) SignOutUserById(id int) error {
	user := models.User{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	return db.Table("Users").Where(&user).Update("loggedin", 0).Error
}
