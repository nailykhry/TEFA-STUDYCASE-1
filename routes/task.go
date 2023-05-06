package routes

import (
	"TEFA-STUDYCASE-1/controllers"
	"TEFA-STUDYCASE-1/middleware"

	"github.com/gofiber/fiber/v2"
)

type TaskRoutes interface {
	Task(app *fiber.App)
}

type taskRoutes struct {
	tasksController controllers.TasksController
}

func NewTaskRoutes(tasksController controllers.TasksController) TaskRoutes {
	return &taskRoutes{tasksController}
}

func (r *taskRoutes) Task(app *fiber.App) {
	app.Post("/tasks", middleware.AuthRequired, r.tasksController.CreateTask)
	app.Get("/tasks", middleware.AuthRequired, r.tasksController.GetTasks)
	app.Get("/tasks/:id", middleware.AuthRequired, r.tasksController.GetTask)
	app.Put("/tasks/:id", middleware.AuthRequired, r.tasksController.UpdateTask)
	app.Delete("/tasks/:id", middleware.AuthRequired, r.tasksController.DeleteTask)
}
