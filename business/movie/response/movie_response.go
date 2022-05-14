package response

import (
	"github.com/mashbens/my-movie-list/business/movie/entity"
	_user "github.com/mashbens/my-movie-list/business/user/response"
)

type MovieResponse struct {
	ID         int                `json:"id"`
	MovieID    string             `json:"movie_id"`
	Title      string             `json:"title"`
	Year       string             `json:"year"`
	Runtime    string             `json:"runtime"`
	Released   string             `json:"released"`
	Genre      string             `json:"genre"`
	Director   string             `json:"director"`
	Writer     string             `json:"writer"`
	Actors     string             `json:"actors"`
	Plot       string             `json:"plot"`
	Language   string             `json:"language"`
	Country    string             `json:"country"`
	Awards     string             `json:"awards"`
	Poster     string             `json:"poster"`
	ImdbRating string             `json:"imdbRating"`
	User       _user.UserResponse `json:"user,omitempty"`
}

func NewMovieResponse(movie entity.Movie) MovieResponse {
	return MovieResponse{
		ID:         movie.ID,
		MovieID:    movie.MovieID,
		Title:      movie.Title,
		Year:       movie.Year,
		Runtime:    movie.Runtime,
		Released:   movie.Released,
		Genre:      movie.Genre,
		Director:   movie.Director,
		Writer:     movie.Writer,
		Actors:     movie.Actors,
		Plot:       movie.Plot,
		Language:   movie.Language,
		Country:    movie.Country,
		Awards:     movie.Awards,
		Poster:     movie.Poster,
		ImdbRating: movie.ImdbRating,
		User:       _user.NewUserResponse(movie.User),
	}
}

func NewMovieArrayResponse(movie []entity.Movie) []MovieResponse {
	movieRes := []MovieResponse{}
	for _, v := range movie {
		m := MovieResponse{
			ID:         v.ID,
			MovieID:    v.MovieID,
			Title:      v.Title,
			Year:       v.Year,
			Runtime:    v.Runtime,
			Released:   v.Released,
			Genre:      v.Genre,
			Director:   v.Director,
			Writer:     v.Writer,
			Actors:     v.Actors,
			Plot:       v.Plot,
			Language:   v.Language,
			Country:    v.Country,
			Awards:     v.Awards,
			Poster:     v.Poster,
			ImdbRating: v.ImdbRating,
			User:       _user.NewUserResponse(v.User),
		}
		movieRes = append(movieRes, m)
	}
	return movieRes
}
