package mysql

import (
	"context"
	"database/sql"

	"bitbucket.com/hexa/common/models"
	"bitbucket.com/hexa/movie"
)

type movieRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewMysqlMovieRepository(db *sql.DB, ctx context.Context) movie.MovieRepository {
	return &movieRepository{
		db:  db,
		ctx: ctx,
	}
}

func (m *movieRepository) GetAllMovies() (movies []*models.Movie, err error) {
	stmt, err := m.db.PrepareContext(m.ctx, "SELECT time_zone FROM cities where coverage_store_id = ? limit 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var moviesP models.Movie
	err = stmt.QueryRowContext(m.ctx).Scan(&moviesP)
	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	default:
		return movies, nil
	}
}

// func (m *movieRepository) GetAllMovies() (movies []*models.Movie, err error) {

// }
