package routes

import (
	"TEFA-STUDYCASE-1/controllers"

	"github.com/gofiber/fiber/v2"
)

type UserchatRoutes interface {
	Userchat(app *fiber.App)
}

type userchatRoutes struct {
	userchatsController controllers.UserchatsController
}

func NewUserchatRoutes(userchatsController controllers.UserchatsController) UserchatRoutes {
	return &userchatRoutes{userchatsController}
}

func (r *userchatRoutes) Userchat(app *fiber.App) {
	app.Post("/chats", r.userchatsController.CreateUserchat)
	app.Get("/chats", r.userchatsController.GetUserchats)
	app.Get("/chats/:id", r.userchatsController.GetUserchat)
	app.Put("/chats/:id", r.userchatsController.UpdateUserchat)
}
