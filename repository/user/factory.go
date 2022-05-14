package user

import (
	"github.com/mashbens/my-movie-list/business/user"
	"github.com/mashbens/my-movie-list/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) user.UserRepository {
	var userRepository user.UserRepository

	if dbCon.Driver == util.MySQL {
		userRepository = NewMysqlRepository(dbCon.MySQL)
	} else if dbCon.Driver == util.PostgreSQL {
		userRepository = NewPostgresRepository(dbCon.PostgreSQL)
	} else {
		panic("Database driver not supported")
	}

	return userRepository
}
