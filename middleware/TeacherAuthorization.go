package middleware

import (
	"TEFA-STUDYCASE-1/security" //O(1)
	"TEFA-STUDYCASE-1/util"     //O(1)

	"github.com/form3tech-oss/jwt-go" //O(1)
	"github.com/gofiber/fiber/v2"     //O(1)
)

func TeacherMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*jwt.Token)       //O(1)
	payload, err := security.ParseToken(token.Raw) //O(n)
	if err != nil {                                //O(1)
		return err //O(1)
	}
	if payload.Subject != "teacher" { //O(1)
		return util.ErrUnauthorized //O(1)
	}
	return ctx.Next() //O(1)
}
