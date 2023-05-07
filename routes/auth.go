package routes

import (
	"TEFA-STUDYCASE-1/controllers"
	"TEFA-STUDYCASE-1/middleware"

	"github.com/gofiber/fiber/v2"
)

type Routes interface {
	Install(app *fiber.App)
}

type authRoutes struct {
	authController controllers.AuthController
}

func NewAuthRoutes(authController controllers.AuthController) Routes {
	return &authRoutes{authController}
}

func (r *authRoutes) Install(app *fiber.App) {
	app.Post("/signup", r.authController.SignUp)
	app.Post("/signin", r.authController.SignIn)
	app.Get("/users", middleware.AuthRequired, middleware.AdminMiddleware, r.authController.GetUsers)
	app.Get("/users/:id", middleware.AuthRequired, r.authController.GetUser)
	app.Put("/users/:id", middleware.AuthRequired, r.authController.PutUser)
	app.Delete("/users/:id", middleware.AuthRequired, middleware.AdminMiddleware, r.authController.DeleteUser)
}
