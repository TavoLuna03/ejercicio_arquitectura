package movie

import (
	"net/http"
)

type MovieHandler interface {
	GetAllMovies(w http.ResponseWriter, r *http.Request)
}

type movieHandler struct {
	movieService MovieService
}

func NewMovieHandler(movieService MovieService) MovieHandler {
	return &movieHandler{
		movieService,
	}
}

func (m *movieHandler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	// products, _ := this.productService.FindAllProductsWithMoreData()
	// response, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// _, _ = w.Write(response)
}
