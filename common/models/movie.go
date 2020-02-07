package models

type Movie struct {
	ID               int     `json:"id" db:"id"`
	VoteCount        int     `json:"vote_count" db:"vote_count"`
	VoteAverage      float64 `json:"vote_average" db:"vote_average"`
	Title            string  `json:"title" db:"title"`
	OriginalTitle    string  `json:"original_title" db:"original_title"`
	OriginalLanguage string  `json:"original_language" db:"original_language"`
	Adult            int     `json:"adult" db:"adult"`
	PosterPath       string  `json:"poster_path" db:"poster_path"`
	Overview         string  `json:"overview" db:"overview"`
	ReleaseDate      string  `json:"release_date" db:"release_date"`
	Popularity       float64 `json:"popularity" db:"popularity"`
}
