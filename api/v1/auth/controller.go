package auth

import (
	"net/http"
	"strconv"

	"github.com/mashbens/my-movie-list/api/common/response"
	service "github.com/mashbens/my-movie-list/business/user"
	"github.com/mashbens/my-movie-list/business/user/dto"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService service.AuthService
	jwtService  service.JWTService
	userService service.UserService
}

func NewAuthController(
	authService service.AuthService,
	jwtService service.JWTService,
	userService service.UserService,
) *AuthController {
	return &AuthController{
		authService: authService,
		jwtService:  jwtService,
		userService: userService,
	}
}

// Login godoc
// @Summary Login
// @Description Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} response.Response
// @Router /api/v1/auth/login [post]
func (controller *AuthController) Login(c echo.Context) error {
	var loginRequest dto.LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = controller.authService.VerifyCredential(loginRequest.Email, loginRequest.Password)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", "Invalid credentials", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	user, _ := controller.userService.FindUserByEmail(loginRequest.Email)

	token := controller.jwtService.GenerateToken((strconv.FormatInt(user.ID, 10)))
	user.Token = token
	response := response.BuildResponse(true, "ok", user)
	return c.JSON(http.StatusOK, response)
}

// RegisterHandler godoc
// @Summary Register
// @Description Register user
// @Tags auth
// @Accept json
// @Produce json
// @Param Name body string true "Name"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} response.Response
// @Router /api/v1/auth/register [post]
func (controller *AuthController) RegisterHandler(c echo.Context) error {
	var registerRequest dto.RegisterRequest
	err := c.Bind(&registerRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	user, err := controller.userService.CreateUser(registerRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	token := controller.jwtService.GenerateToken((strconv.FormatInt(user.ID, 10)))
	user.Token = token
	response := response.BuildResponse(true, "ok", user)
	return c.JSON(http.StatusOK, response)
}
