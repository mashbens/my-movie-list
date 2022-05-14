package entity

import (
	"github.com/mashbens/my-movie-list/business/user/entity"
)

type Movie struct {
	ID         int    `gorm:"primary_key:auto_increment" json:"-"`
	MovieID    string `gorm:"type:varchar(100)" json:"-"`
	Title      string `gorm:"type:varchar(100)" json:"-"`
	Year       string `gorm:"type:varchar(100)" json:"-"`
	Runtime    string `gorm:"type:varchar(100)" json:"-"`
	Released   string `gorm:"type:varchar(100)" json:"-"`
	Genre      string `gorm:"type:varchar(100)" json:"-"`
	Director   string `gorm:"type:varchar(100)" json:"-"`
	Writer     string `gorm:"type:varchar(100)" json:"-"`
	Actors     string `gorm:"type:varchar(100)" json:"-"`
	Plot       string `gorm:"type:longtext" json:"-"`
	Language   string `gorm:"type:varchar(100)" json:"-"`
	Country    string `gorm:"type:varchar(100)" json:"-"`
	Awards     string `gorm:"type:varchar(100)" json:"-"`
	Poster     string `gorm:"type:longtext" json:"-"`
	ImdbRating string `gorm:"type:varchar(100)" json:"-"`
	UserID     int64
	User       entity.User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
