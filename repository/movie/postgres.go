package movie

import (
	"github.com/mashbens/my-movie-list/business/movie"
	"github.com/mashbens/my-movie-list/business/movie/entity"

	"gorm.io/gorm"
)

type PostgreMovieRepository struct {
	db *gorm.DB
}

func NewPosgresMovieRepository(db *gorm.DB) movie.MovieRepository {
	return &PostgreMovieRepository{
		db: db,
	}
}

func (c *PostgreMovieRepository) All(userID string) ([]entity.Movie, error) {
	movies := []entity.Movie{}

	c.db.Preload("User").Where("user_id = ?", userID).Find(&movies)
	return movies, nil
}

func (c *PostgreMovieRepository) InsertMovie(movie entity.Movie) (entity.Movie, error) {
	c.db.Save(&movie)
	c.db.Preload("User").Find(&movie)
	return movie, nil
}

func (c *PostgreMovieRepository) FindOneMovieByID(ID string) (entity.Movie, error) {
	var movie entity.Movie
	res := c.db.Preload("User").Where("id = ?", ID).First(&movie)
	if res.Error != nil {
		return movie, res.Error
	}
	return movie, nil
}

func (c *PostgreMovieRepository) DelMovie(movieID string) error {
	var movie entity.Movie
	res := c.db.Preload("User").Where("id = ?", movieID).First(&movie)
	if res.Error != nil {
		return res.Error
	}
	c.db.Delete(&movie)
	return nil
}
