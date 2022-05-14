package user

import (
	"github.com/mashbens/my-movie-list/business/user/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	col *mongo.Collection
}

type collection struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"first_name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func newCollection(user entity.User) collection {
	return collection{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

// belom ditest
