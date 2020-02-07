package movie

import (
	"bitbucket.com/hexa/common/models"
)

// ProductService es una representación de un
type MovieService interface {
	GetAllMovies() ([]*models.Movie, error)
}

type movieService struct {
	repoMovie MovieRepository
}

func NewMoviesService(repoMovie MovieRepository) MovieService {
	return &movieService{
		repoMovie,
	}
}

func (this *movieService) GetAllMovies() (movies []*models.Movie, err error) {
	movies, err = this.repoMovie.GetAllMovies()
	if err != nil {
		return nil, err
	}
	return movies, nil
}
