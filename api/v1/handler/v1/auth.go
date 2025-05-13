package v1

import (
	"github.com/achmad-dev/simple-ewallet/internal/dto"
	"github.com/achmad-dev/simple-ewallet/internal/pkg"
	"github.com/achmad-dev/simple-ewallet/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Signup(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	userService service.UserService
}

// Login implements AuthHandler.
func (a *authHandler) Login(c *fiber.Ctx) error {
	var userLoginDto dto.AuthDto
	if err := c.BodyParser(&userLoginDto); err != nil {
		return pkg.ErrorResponse(c, fiber.StatusBadRequest, "invalid request")
	}

	authResponse, err := a.userService.SignIn(c.Context(), userLoginDto.Username, userLoginDto.Password)
	if err != nil {
		return pkg.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.SuccessResponse(c, authResponse, "user logged in successfully")
}

// Signup implements AuthHandler.
func (a *authHandler) Signup(c *fiber.Ctx) error {
	var userSignupDto dto.UserSignupDto
	if err := c.BodyParser(&userSignupDto); err != nil {
		return pkg.ErrorResponse(c, fiber.StatusBadRequest, "invalid request")
	}

	authResponse, err := a.userService.Signup(c.Context(), userSignupDto)
	if err != nil {
		return pkg.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.SuccessResponse(c, authResponse, "user created successfully")
}

func NewAuthHandler(userService service.UserService) AuthHandler {
	return &authHandler{
		userService: userService,
	}
}
