package middleware

import (
	"TEFA-STUDYCASE-1/security"
	"TEFA-STUDYCASE-1/util"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		return err
	}
	if payload.Subject != "admin" {
		return util.ErrUnauthorized
	}
	return ctx.Next()
}
