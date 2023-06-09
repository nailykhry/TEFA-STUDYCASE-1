package routes

import (
	"TEFA-STUDYCASE-1/controllers"
	"TEFA-STUDYCASE-1/middleware"

	"github.com/gofiber/fiber/v2"
)

type UsersubRoutes interface {
	Usersub(app *fiber.App)
}

type usersubRoutes struct {
	usersubsController controllers.UsersubsController
}

func NewUsersubRoutes(usersubsController controllers.UsersubsController) UsersubRoutes {
	return &usersubRoutes{usersubsController}
}

func (r *usersubRoutes) Usersub(app *fiber.App) {
	app.Post("/usersubs", middleware.AuthRequired, r.usersubsController.CreateUsersub)
}
