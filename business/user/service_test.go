package user_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/mashbens/my-movie-list/business/user"

	// movieEntity "github.com/mashbens/my-movie-list/business/movie/entity"
	userEntity "github.com/mashbens/my-movie-list/business/user/entity"
	"github.com/mashbens/my-movie-list/business/user/response"
)

var service user.UserService
var user1, user2 userEntity.User

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindUserByID(t *testing.T) {
	t.Run("Expect to find user by id", func(t *testing.T) {
		userID := int(user1.ID)
		userIDs := strconv.Itoa(userID)
		user, err := service.FindUserByID(userIDs)
		if err != nil {
			t.Error("Expect to find user by id", err)
		} else {
			if user.ID != 1 {
				t.Errorf("Expected %d, got %d", 0, user.ID)
			}
		}
	})
}

func TestFindUserByEmail(t *testing.T) {
	t.Run("Expect to find user by email", func(t *testing.T) {
		user, err := service.FindUserByEmail(user1.Email)
		if err != nil {
			t.Error("Expect to find user by email", err)
		} else {
			if user.ID != 1 {
				t.Errorf("Expected %d, got %d", 0, user.ID)
			}
		}
	})
}

func TestInsertUser(t *testing.T) {

}

func TestUpdateUser(t *testing.T) {

}

type inMemoryRepository struct {
	userByID    map[string]userEntity.User
	UserByEmail map[string]userEntity.User

	userResponse map[string]response.UserResponse
}

func newInMemoryRepository() *inMemoryRepository {
	var repo inMemoryRepository
	repo.userByID = make(map[string]userEntity.User)
	repo.UserByEmail = make(map[string]userEntity.User)

	userID := int64(user1.ID)
	userIDs := strconv.FormatInt(userID, 10)
	repo.userByID[userIDs] = user1
	repo.UserByEmail[user1.Email] = user1

	return &repo
}
func (r *inMemoryRepository) InsertUser(userEntity.User) (userEntity.User, error) {
	return userEntity.User{}, nil
}

func (r *inMemoryRepository) FindByEmail(email string) (userEntity.User, error) {
	return r.UserByEmail[email], nil
}

func (r *inMemoryRepository) FindByUserID(userID string) (userEntity.User, error) {
	return r.userByID[userID], nil
}

func (r *inMemoryRepository) UpdateUser(user userEntity.User) (userEntity.User, error) {
	return user, nil
}

func setup() {
	user1.ID = 1
	user1.Name = "John"
	user1.Email = "test@mail.mail"
	user1.Password = "test123"

	user2.ID = 2
	user2.Name = "Jane"
	user2.Email = "Jane@mail.mail"
	user2.Password = "test123"

	repo := newInMemoryRepository()
	service = user.NewUserService(repo)
}
