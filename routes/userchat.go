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
	app.Post("/tasks", r.userchatsController.CreateUserchat)
	app.Get("/tasks", r.userchatsController.GetUserchats)
	app.Get("/tasks/:id", r.userchatsController.GetUserchat)
	app.Put("/tasks/:id", r.userchatsController.UpdateUserchat)
}
