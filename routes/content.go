package routes

import (
	"TEFA-STUDYCASE-1/controllers"
	"TEFA-STUDYCASE-1/middleware"

	"github.com/gofiber/fiber/v2"
)

type ContentRoutes interface {
	Content(app *fiber.App)
}

type contentsRoutes struct {
	contentsController controllers.ContentsController
}

func NewContentRoutes(contentsController controllers.ContentsController) ContentRoutes {
	return &contentsRoutes{contentsController}
}

func (r *contentsRoutes) Content(app *fiber.App) {
	app.Post("/contents", middleware.AuthRequired, r.contentsController.UploadContent)
	app.Get("/contents", middleware.AuthRequired, r.contentsController.GetContents)
	app.Get("/contents/:id", middleware.AuthRequired, r.contentsController.GetContent)
	app.Put("/contents/:id", middleware.AuthRequired, r.contentsController.PutContent)
	app.Delete("/contents/:id", middleware.AuthRequired, middleware.AdminMiddleware, r.contentsController.DeleteContent)
}
