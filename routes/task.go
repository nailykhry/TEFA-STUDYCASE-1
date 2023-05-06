package routes

import (
	"TEFA-STUDYCASE-1/controllers"

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
	app.Post("/tasks", r.tasksController.CreateTask)
	app.Get("/tasks", r.tasksController.GetTasks)
	app.Get("/tasks/:id", r.tasksController.GetTask)
	app.Put("/tasks/:id", r.tasksController.UpdateTask)
	app.Delete("/tasks/:id", r.tasksController.DeleteTask)
}
