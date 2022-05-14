package user

import (
	"log"

	"github.com/mashbens/my-movie-list/business/user"
	"github.com/mashbens/my-movie-list/business/user/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) user.UserRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (c *MysqlRepository) InsertUser(user entity.User) (entity.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	c.db.Save(&user)
	return user, nil
}

func (c *MysqlRepository) UpdateUser(user entity.User) (entity.User, error) {
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

func (c *MysqlRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	res := c.db.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (c *MysqlRepository) FindByUserID(userID string) (entity.User, error) {
	var user entity.User
	res := c.db.Where("id = ?", userID).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
