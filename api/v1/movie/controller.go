package movie

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mashbens/my-movie-list/api/common/obj"
	"github.com/mashbens/my-movie-list/api/common/response"
	movieService "github.com/mashbens/my-movie-list/business/movie"
	"github.com/mashbens/my-movie-list/business/movie/dto"
	jwtService "github.com/mashbens/my-movie-list/business/user"

	"github.com/eefret/gomdb"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type MovieController struct {
	movieService movieService.MovieService
	jwtService   jwtService.JWTService
}

func NewMovieController(
	movieService movieService.MovieService,
	jwtService jwtService.JWTService,
) *MovieController {
	return &MovieController{
		movieService: movieService,
		jwtService:   jwtService,
	}
}

// All godoc
// @Summary Get all user movies
// @Description Get all user movies(watchlist),  Header[Authorization]: Token
// @Tags movie
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/movie/all [get]
func (controller *MovieController) All(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(authHeader, c)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	movies, err := controller.movieService.All(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("Failed to process request", "Failed to get movies", nil))
	}
	response := movies
	return c.JSON(http.StatusOK, response)
}

// AddWishlist godoc
// @Summary Add movie to watchlist
// @Description Add movie to watchlist, Header[Authorization]: Token
// @Tags movie
// @Accept json
// @Produce json
// @Param ID body string true "Movie ID"
// @Success 200 {object} response.Response
// @Router /api/v1/movie/add [post]
func (controller *MovieController) AddWishList(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(authHeader, c)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	req := new(dto.MovieRequest)
	c.Bind(req)
	id := req.ID

	api := gomdb.Init("11ab3263")
	res, err := api.MovieByImdbID(id)
	if err != nil {
		return err
	}
	intuserID, _ := strconv.Atoi(userID)
	movie := dto.CreateMovieRequest{
		MovieID:  req.ID,
		Title:    res.Title,
		Year:     res.Year,
		Runtime:  res.Runtime,
		Released: res.Released,
		Genre:    res.Genre,
		Director: res.Director,
		Writer:   res.Writer,
		Actors:   res.Actors,
		Plot:     res.Plot,
		Language: res.Language,
		Country:  res.Country,
		Awards:   res.Awards,
		Poster:   res.Poster,
		UserID:   int64(intuserID),
	}

	movies, err := controller.movieService.AddMovie(movie, int64(intuserID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("Failed to process request", "Failed to get movies", nil))
	}
	response := movies
	return c.JSON(http.StatusOK, response)
}

// SearchMovie godoc
// @Summary Search movie
// @Description Search movie, Header[Authorization]: Token, Param: movie name
// @Description Example: /api/v1/movie/search/avengers
// @Tags movie
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/v1/movie/search [post]
func (controller *MovieController) SearchMovie(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(authHeader, c)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	_ = userID
	req := new(dto.MovieRequest)
	c.Bind(req)
	api := gomdb.Init("11ab3263")
	query := &gomdb.QueryData{Title: req.Search, SearchType: gomdb.MovieSearch}
	res, err := api.Search(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response := res
	return c.JSON(http.StatusOK, response)
}

// FindOneMovieByID godoc
// @Summary Find movie by ID From watchlist
// @Description Find movie by ID From watchlist, Header[Authorization]: Token
// @Description Example: /api/v1/movie/mylist/1
// @Tags movie
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/movie/mylist/:id [get]
func (controller *MovieController) FindOneMovieByID(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(authHeader, c)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	_ = userID

	req := new(dto.MovieRequest)
	c.Bind(req)

	reqId := req.ID
	movie, err := controller.movieService.FindMovieByID(reqId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("Failed to process request", "Failed to get movies", nil))
	}
	response := movie
	return c.JSON(http.StatusOK, response)
}

// DeleteMovieFromWatchlist godoc
// @Summary Delete movie from watchlist
// @Description Delete movie from watchlist, Header[Authorization]: Token
// @Description Example: /api/v1/movie/mylist/1
// @Tags movie
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/movie/mylist/:id [delete]
func (controller *MovieController) DeleteMovie(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(authHeader, c)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	req := new(dto.MovieRequest)
	c.Bind(req)

	reqId := req.ID

	err := controller.movieService.DeleteMovie(reqId, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("Failed to process request", "Failed to get movies", nil))
	}

	response := response.BuildResponse(true, "Movie deleted", obj.EmptyObj{})
	return c.JSON(http.StatusOK, response)
}
