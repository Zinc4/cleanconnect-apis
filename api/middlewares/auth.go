package middlewares

import (
	"errors"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte("secret")},
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			authHeader := c.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Missing or invalid token",
				})
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte("secret"), nil
			})

			if err != nil || !token.Valid {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid token",
				})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid token claims",
				})
			}

			userID := uint(claims["userID"].(float64))
			c.Locals("userID", userID)

			return c.Next()
		},
	})

}

func AdminOnly() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte("secret")},
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			claims, _ := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
			if claims["role"] != "admin" {
				return c.Status(fiber.StatusForbidden).
					JSON(fiber.Map{"status": "error", "message": "Forbidden", "data": nil})
			}
			return c.Next()
		},
	})
}

func SuperAdminOnly() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte("secret")},
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			claims, _ := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
			if claims["role"] != "admin" && claims["role"] != "superadmin" {
				return c.Status(fiber.StatusForbidden).
					JSON(fiber.Map{"status": "error", "message": "Forbidden", "data": nil})
			}
			return c.Next()
		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
