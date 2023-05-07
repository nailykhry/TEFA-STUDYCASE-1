package middleware

import (
	"TEFA-STUDYCASE-1/security" //O(1)
	"TEFA-STUDYCASE-1/util"     //O(1)
	"net/http"                  //O(1)

	"github.com/gofiber/fiber/v2"       //O(1)
	jwtware "github.com/gofiber/jwt/v2" //O(1)
)

func AuthRequired(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{ //O(1)
		SigningKey:    security.JwtSecretKey,
		SigningMethod: security.JwtSigningMethod,
		TokenLookup:   "header:Authorization",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.
				Status(http.StatusUnauthorized).
				JSON(util.NewJError(err))
		},
	})(ctx)
}
