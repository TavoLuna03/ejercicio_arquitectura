package movie

import (
	"bitbucket.com/hexa/common/models"
)

type MovieRepository interface {
	GetAllMovies() ([]*models.Movie, error)
}
