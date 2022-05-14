package movie

import (
	"github.com/mashbens/my-movie-list/business/movie"
	"github.com/mashbens/my-movie-list/business/movie/entity"

	"gorm.io/gorm"
)

type MysqlmovieRepository struct {
	db *gorm.DB
}

func NewMysqlMovieRepository(db *gorm.DB) movie.MovieRepository {
	return &MysqlmovieRepository{
		db: db,
	}
}

func (c *MysqlmovieRepository) All(userID string) ([]entity.Movie, error) {
	movies := []entity.Movie{}

	c.db.Preload("User").Where("user_id = ?", userID).Find(&movies)
	return movies, nil
}

func (c *MysqlmovieRepository) InsertMovie(movie entity.Movie) (entity.Movie, error) {
	c.db.Create(&movie)
	c.db.Preload("User").Find(&movie)
	return movie, nil
}

func (c *MysqlmovieRepository) FindOneMovieByID(ID string) (entity.Movie, error) {
	var movie entity.Movie
	res := c.db.Preload("User").Where("id = ?", ID).First(&movie)
	if res.Error != nil {
		return movie, res.Error
	}
	return movie, nil
}

func (c *MysqlmovieRepository) DelMovie(movieID string) error {
	var movie entity.Movie
	res := c.db.Preload("User").Where("id = ?", movieID).First(&movie)
	if res.Error != nil {
		return res.Error
	}
	c.db.Delete(&movie)
	return nil
}
