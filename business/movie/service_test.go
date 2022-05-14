package movie_test

import (
	"os"

	"github.com/mashbens/my-movie-list/business/movie"
	movieEntity "github.com/mashbens/my-movie-list/business/movie/entity"
	userEntity "github.com/mashbens/my-movie-list/business/user/entity"

	"strconv"
	"testing"
)

var service movie.MovieService
var movie1, movie2 movieEntity.Movie

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetAllMovies(t *testing.T) {
	t.Run("Expect to get all movies", func(t *testing.T) {

		foundMovie, _ := service.All(strconv.Itoa(movie1.ID))
		if len(*foundMovie) != 1 {
			t.Errorf("Expected to get 1 movie, got %d", len(*foundMovie))
		}
	})
}

func TestGetMovieByID(t *testing.T) {
	t.Run("Expect to get movie by id", func(t *testing.T) {
		movie, err := service.FindMovieByID(strconv.Itoa(movie1.ID))
		if err != nil {
			t.Error("Expect to get movie by id", err)
		} else {
			if movie.ID != 0 {
				t.Errorf("Expected %d, got %d", 0, movie.ID)
			}
		}
	})
}

type inMemoryRepository struct {
	AllMovie map[string]movieEntity.Movie
}

func newInMemoryRepository() *inMemoryRepository {
	var repo inMemoryRepository
	repo.AllMovie = make(map[string]movieEntity.Movie)

	repo.AllMovie[movie1.MovieID] = movie1

	return &repo
}

func (r *inMemoryRepository) All(userID string) ([]movieEntity.Movie, error) {
	var movies []movieEntity.Movie
	for _, movie := range r.AllMovie {
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r *inMemoryRepository) InsertMovie(movie movieEntity.Movie) (movieEntity.Movie, error) {
	r.AllMovie[movie.MovieID] = movie
	return movie, nil
}

func (r *inMemoryRepository) FindOneMovieByID(ID string) (movieEntity.Movie, error) {
	return r.AllMovie[ID], nil
}

func (r *inMemoryRepository) DelMovie(movieID string) error {
	delete(r.AllMovie, movieID)
	return nil
}

func setup() {
	movie1.ID = 3
	movie1.MovieID = "tt0111161"
	movie1.Title = "The Shawshank Redemption"
	movie1.Year = "1994"
	movie1.Runtime = "142 min"
	movie1.Released = "14-10-10"
	movie1.Genre = "Drama"
	movie1.Director = "Frank Darabont"
	movie1.Writer = "Stephen King (short story), Frank Darabont (screenplay)"
	movie1.Actors = "Tim Robbins, Morgan Freeman, Bob Gunton, William Sadler"
	movie1.Plot = "Chronicles the experiences of a formerly successful banker as a prisoner in the gloomy jailhouse of Shawshank after being found guilty of a crime he did not commit. The film portrays the man's unique way of dealing with his new, torturous life; along the way he befriends a number of fellow prisoners, most notably a wise long-term inmate named Red."
	movie1.Language = "English"
	movie1.Country = "USA"
	movie1.Awards = "Nominated for 7 Oscars. 21 wins & 43 nominations total"
	movie1.Poster = "https://m.media-amazon.com/images/M/MV5BODU4MjU4NjIwNl5BMl5BanBnXkFtZTgwMDU2MjEyMDE@._V1_SX300.jpg"
	movie1.ImdbRating = "80"
	movie1.UserID = 1
	movie1.User = userEntity.User{
		ID:       1,
		Name:     "admin",
		Password: "admin",
	}

	movie2.MovieID = "tt0111161"
	movie2.Title = "The Shawshank Redemption"
	movie2.Year = "1994"
	movie2.Runtime = "142 min"
	movie2.Released = "14-10-10"
	movie2.Genre = "Drama"
	movie2.Director = "Frank Darabont"
	movie2.Writer = "Stephen King (short story), Frank Darabont (screenplay)"
	movie2.Actors = "Tim Robbins, Morgan Freeman, Bob Gunton, William Sadler"
	movie2.Plot = "Chronicles the experiences of a formerly successful banker as a prisoner in the gloomy jailhouse of Shawshank after being found guilty of a crime he did not commit. The film portrays the man's unique way of dealing with his new, torturous life; along the way he befriends a number of fellow prisoners, most notably a wise long-term inmate named Red."
	movie2.Language = "English"
	movie2.Country = "USA"
	movie2.Awards = "Nominated for 7 Oscars. 21 wins & 43 nominations total"
	movie2.Poster = "https://m.media-amazon.com/images/M/MV5BODU4MjU4NjIwNl5BMl5BanBnXkFtZTgwMDU2MjEyMDE@._V1_SX300.jpg"
	movie2.ImdbRating = "80"
	movie2.UserID = 1
	movie2.User = userEntity.User{
		ID:       1,
		Name:     "admin",
		Password: "admin",
	}

	repo := newInMemoryRepository()
	service = movie.NewMovieService(repo)
}
