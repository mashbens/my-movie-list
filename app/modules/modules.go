package modules

import (
	"github.com/mashbens/my-movie-list/api"
	"github.com/mashbens/my-movie-list/api/v1/auth"
	"github.com/mashbens/my-movie-list/api/v1/movie"
	"github.com/mashbens/my-movie-list/api/v1/user"
	movieService "github.com/mashbens/my-movie-list/business/movie"
	authService "github.com/mashbens/my-movie-list/business/user"
	jwtService "github.com/mashbens/my-movie-list/business/user"
	userService "github.com/mashbens/my-movie-list/business/user"
	"github.com/mashbens/my-movie-list/config"
	movieRepo "github.com/mashbens/my-movie-list/repository/movie"
	userRepo "github.com/mashbens/my-movie-list/repository/user"
	"github.com/mashbens/my-movie-list/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	userRepo := userRepo.RepositoryFactory(dbCon)

	userService := userService.NewUserService(userRepo)
	authService := authService.NewAuthService(userRepo)
	jwtService := jwtService.NewJWTService()

	movieRepo := movieRepo.MovieRepositoryFactory(dbCon)
	movieService := movieService.NewMovieService(movieRepo)

	controller := api.Controller{
		Auth:  auth.NewAuthController(authService, jwtService, userService),
		User:  user.NewUserController(userService, jwtService),
		Movie: movie.NewMovieController(movieService, jwtService),
	}

	return controller
}
