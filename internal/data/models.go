package data

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Create a Models struct which wraps the MovieModel. We'll add other models to this,
// like a UserModel and PermissionModel, as our build progresses.
type Models struct {
	Movies interface {
		Insert(movie *Movie, r *http.Request) error
		Get(id int64, r *http.Request) (*Movie, error)
		Update(movie *Movie, r *http.Request) error
		Delete(id int64, r *http.Request) error
		GetAll(title string, genres []string, filters Filters, r *http.Request) ([]*Movie, Metadata, error)
	}
	Users interface {
		Insert(user *User, r *http.Request) error
		GetByEmail(email string, r *http.Request) (*User, error)
		Update(user *User, r *http.Request) error
	}
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized MovieModel.
func NewModels(db *pgxpool.Pool) Models {
	m := MovieModel{DB: db}
	u := UserModel{
		DB: db,
	}
	return Models{
		Movies: m,
		Users:  u,
	}
}

func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
		Users:  MockUserModel{},
	}
}
