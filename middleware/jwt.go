package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/model"
)

func AuthenticateJWT(requireAdmin bool, config configuration.Config) func(*fiber.Ctx) error {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	return jwtware.New(jwtware.Config{
		SigningKey:  []byte(jwtSecret),
		TokenLookup: "header:token",
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			userIDFloat := claims["userID"].(float64)
			userID := int(userIDFloat)
			noTelp := claims["noTelp"].(string)
			isAdmin := claims["is_admin"].(bool)

			ctx.Locals("userID", userID)
			ctx.Locals("noTelp", noTelp)
			ctx.Locals("is_admin", isAdmin)

			if requireAdmin && !isAdmin {
				return ctx.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
					Status:  false,
					Message: "Failed to POST data",
					Errors:  []string{"Unauthorized"},
					Data:    nil,
				})
			}

			return ctx.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
					Status:  false,
					Message: "Failed to POST data",
					Errors:  []string{"Missing or malformed JWT"},
					Data:    nil,
				})
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
					Status:  false,
					Message: "Failed to POST data",
					Errors:  []string{"Expired or invalid JWT"},
					Data:    nil,
				})
			}
		},
	})

}
