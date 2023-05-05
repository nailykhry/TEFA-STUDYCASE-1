package routes

import (
	"TEFA-STUDYCASE-1/controllers"

	"github.com/gofiber/fiber/v2"
)

type Routes interface {
	Content(app *fiber.App)
}

type contentsRoutes struct {
	contentsController controllers.ContentsController
}

func NewContentRoutes(contentsController controllers.ContentsController) Routes {
	return &contentsRoutes{contentsController}
}

func (r *contentsRoutes) Content(app *fiber.App) {
	app.Post("/contents", r.contentsController.UploadContent)
	app.Get("/contents", r.contentsController.GetContents)
	app.Get("/contents/:id", r.contentsController.GetContent)
	app.Put("/contents/:id", r.contentsController.PutContent)
	app.Delete("/contents/:id", r.contentsController.DeleteContent)
}
