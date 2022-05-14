package movie

import (
	// "github.com/mashbens/my-movie-list/business/movie/dto"

	"fmt"

	"github.com/mashbens/my-movie-list/business/movie/dto"
	"github.com/mashbens/my-movie-list/business/movie/entity"
	_movie "github.com/mashbens/my-movie-list/business/movie/response"

	"github.com/mashingan/smapping"
)

type MovieRepository interface {
	All(userID string) ([]entity.Movie, error)
	InsertMovie(movie entity.Movie) (entity.Movie, error)
	FindOneMovieByID(ID string) (entity.Movie, error)
	DelMovie(movieID string) error
}

type MovieService interface {
	All(userID string) (*[]_movie.MovieResponse, error)
	AddMovie(addMovie dto.CreateMovieRequest, userID int64) (_movie.MovieResponse, error)
	DeleteMovie(movieID string, userID string) error
	FindMovieByID(movieID string) (_movie.MovieResponse, error)
}

type movieService struct {
	movieRepo MovieRepository
}

func NewMovieService(movieRepo MovieRepository) MovieService {
	return &movieService{
		movieRepo: movieRepo,
	}
}

func (c *movieService) All(userID string) (*[]_movie.MovieResponse, error) {
	movies, err := c.movieRepo.All(userID)
	if err != nil {
		return nil, err
	}
	movs := _movie.NewMovieArrayResponse(movies)
	return &movs, nil
}

func (c *movieService) AddMovie(addMovie dto.CreateMovieRequest, userID int64) (_movie.MovieResponse, error) {

	movie := entity.Movie{}
	err := smapping.FillStruct(&movie, smapping.MapFields(&addMovie))

	if err != nil {
		return _movie.MovieResponse{}, err
	}
	id := userID
	movie.UserID = id
	m, err := c.movieRepo.InsertMovie(movie)
	if err != nil {
		return _movie.MovieResponse{}, err
	}
	res := _movie.NewMovieResponse(m)
	return res, nil
}

func (c *movieService) FindMovieByID(movieID string) (_movie.MovieResponse, error) {
	movie, err := c.movieRepo.FindOneMovieByID(movieID)
	if err != nil {
		return _movie.MovieResponse{}, err
	}
	res := _movie.NewMovieResponse(movie)
	return res, nil
}

func (c *movieService) DeleteMovie(movieID string, userID string) error {
	movie, err := c.movieRepo.FindOneMovieByID(movieID)
	if err != nil {
		return err
	}
	if fmt.Sprintf("%d", movie.UserID) != userID {
		return fmt.Errorf("user not match")

	}
	c.movieRepo.DelMovie(movieID)
	return nil
}
