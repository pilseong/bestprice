package gmarket

import (
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllItems(pageNo, pageSize int) ([]*Item, error)
	// GetUserByEmail(email string) (*models.User, error)
	// GetUserByID(id int) (*models.User, error)

	// OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)
	// OneMovie(id int) (*models.Movie, error)
	// AllGenres() ([]*models.Genre, error)
	// InsertMovie(movie models.Movie) (int, error)
	// UpdateMovieGenres(id int, genreIDs []int) error
	// UpdateMovie(movie models.Movie) error
	// DeleteMovie(id int) error
}
