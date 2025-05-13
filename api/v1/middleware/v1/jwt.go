package v1

import (
	"strings"

	"github.com/achmad-dev/simple-ewallet/internal/pkg"
	"github.com/achmad-dev/simple-ewallet/internal/service"
	"github.com/gofiber/fiber/v2"
)

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/

func AuthMiddleware(secret string, userService service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
		}
		claims, err := pkg.ValidateToken(tokenString, secret)
		if err != nil {
			return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
		}

		user, err := userService.GetUserByID(c.Context(), claims.UserId)
		if err != nil {
			return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
		}
		c.Locals("user_id", string(user.ID))
		return c.Next()
	}
}
