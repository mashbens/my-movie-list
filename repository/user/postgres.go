package user

import (
	// "log"
	"github.com/mashbens/my-movie-list/business/user"
	"github.com/mashbens/my-movie-list/business/user/entity"

	// "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) user.UserRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (c *PostgresRepository) InsertUser(user entity.User) (entity.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	c.db.Save(&user)
	return user, nil
}

func (c *PostgresRepository) UpdateUser(user entity.User) (entity.User, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		c.db.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	c.db.Save(&user)
	return user, nil
}

func (c *PostgresRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	res := c.db.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (c *PostgresRepository) FindByUserID(userID string) (entity.User, error) {
	var user entity.User
	res := c.db.Where("id = ?", userID).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}
