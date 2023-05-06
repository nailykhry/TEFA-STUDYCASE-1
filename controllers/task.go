package controllers

import (
	"TEFA-STUDYCASE-1/models"     // O(1)
	"TEFA-STUDYCASE-1/repository" // O(1)
	"TEFA-STUDYCASE-1/util"       // O(1)
	"errors"                      // O(1)
	"net/http"                    // O(1)
	"time"                        // O(1)

	"github.com/gofiber/fiber/v2" // O(1)
	"gopkg.in/mgo.v2/bson"        // O(1)
)

type TasksController interface {
	CreateTask(ctx *fiber.Ctx) error // O(1)
	GetTask(ctx *fiber.Ctx) error    // O(1)
	GetTasks(ctx *fiber.Ctx) error   // O(1)
	UpdateTask(ctx *fiber.Ctx) error // O(1)
	DeleteTask(ctx *fiber.Ctx) error // O(1)
}

func NewTaskController(tasksRepo repository.TasksRepository) TasksController {
	return &tasksController{tasksRepo} // O(1)
}

type tasksController struct {
	taskRepo repository.TasksRepository // O(1)
}

func (c *tasksController) CreateTask(ctx *fiber.Ctx) error {
	var newTask models.Task         // O(1)
	err := ctx.BodyParser(&newTask) // O(1)
	if err != nil {                 // O(1)
		return ctx. // O(1)
				Status(http.StatusUnprocessableEntity).
				JSON(util.NewJError(err))
	}

	if newTask.Title == "" || newTask.Description == "" { // O(1)
		return ctx. // O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(errors.New("bad request: invalid task")))
	}

	newTask.CreatedAt = time.Now()        // O(1)
	newTask.UpdatedAt = newTask.CreatedAt // O(1)
	newTask.Id = bson.NewObjectId()       // O(1)

	err = c.taskRepo.CreateTask(&newTask) // O(1)
	if err != nil {                       // O(1)
		return ctx. // O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
	}
	return ctx. // O(1)
			Status(http.StatusCreated).
			JSON(newTask)
}

func (c *tasksController) GetTask(ctx *fiber.Ctx) error {
	taskID := ctx.Params("id")       // O(1)
	if !bson.IsObjectIdHex(taskID) { // O(1)
		return ctx.Status(http.StatusBadRequest).
			JSON(util.NewJError(errors.New("invalid task id")))
	}
	user, err := c.taskRepo.GetTaskById(taskID) // O(1)
	if err != nil {                             // O(1)
		return ctx. // O(1)
				Status(http.StatusInternalServerError).
				JSON(util.NewJError(err))
	}
	return ctx. // O(1)
			Status(http.StatusOK).
			JSON(user)
}

func (c *tasksController) GetTasks(ctx *fiber.Ctx) error {
	tasks, err := c.taskRepo.GetAllTasks() // O(n)
	if err != nil {                        // O(1)
		return ctx. // O(1)
				Status(http.StatusInternalServerError).
				JSON(util.NewJError(err))
	}
	return ctx. // O(1)
			Status(http.StatusOK).
			JSON(tasks)
}

func (c *tasksController) UpdateTask(ctx *fiber.Ctx) error {
	taskID := ctx.Params("id")     // O(1)
	var update models.Task         // O(1)
	err := ctx.BodyParser(&update) // O(1)
	if err != nil {                // O(1)
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	update.UpdatedAt = time.Now()        // O(1)
	update.Id = bson.ObjectIdHex(taskID) // O(1)

	err = c.taskRepo.UpdateTask(&update) // O(1)
	if err != nil {                      // O(1)
		return ctx. // O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
	}

	return ctx. // O(1)
			Status(http.StatusOK).
			JSON(update)
}

func (c *tasksController) DeleteTask(ctx *fiber.Ctx) error {
	taskID := ctx.Params("id")       // O(1)
	if !bson.IsObjectIdHex(taskID) { // O(1)
		return ctx.Status(http.StatusBadRequest). // O(1)
								JSON(util.NewJError(errors.New("invalid task id")))
	}

	err := c.taskRepo.DeleteTask(taskID) // O(1)
	if err != nil {                      // O(1)
		return ctx. // O(1)
				Status(http.StatusInternalServerError).
				JSON(util.NewJError(err))
	}
	ctx.Set("Entity", taskID)            // O(1)
	return ctx.SendStatus(http.StatusOK) // O(1)
}
