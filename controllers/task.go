package controllers

import (
	"TEFA-STUDYCASE-1/models"
	"TEFA-STUDYCASE-1/repository"
	"TEFA-STUDYCASE-1/util"
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2/bson"
)

type TasksController interface {
	CreateTask(ctx *fiber.Ctx) error
	GetTask(ctx *fiber.Ctx) error
	GetTasks(ctx *fiber.Ctx) error
	UpdateTask(ctx *fiber.Ctx) error
	DeleteTask(ctx *fiber.Ctx) error
}

func NewTaskController(tasksRepo repository.TasksRepository) TasksController {
	return &tasksController{tasksRepo}
}

type tasksController struct {
	taskRepo repository.TasksRepository
}

func (c *tasksController) CreateTask(ctx *fiber.Ctx) error {
	var newTask models.Task
	err := ctx.BodyParser(&newTask)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	if newTask.Title == "" || newTask.Description == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(errors.New("bad request: invalid task")))
	}

	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = newTask.CreatedAt
	newTask.Id = bson.NewObjectId()

	err = c.taskRepo.CreateTask(&newTask)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusCreated).
		JSON(newTask)
}

func (c *tasksController) GetTask(ctx *fiber.Ctx) error {
	taskID := ctx.Params("id")
	if !bson.IsObjectIdHex(taskID) {
		return ctx.Status(http.StatusBadRequest).
			JSON(util.NewJError(errors.New("invalid content id")))
	}
	user, err := c.taskRepo.GetTaskById(taskID)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(user)
}

func (c *tasksController) GetTasks(ctx *fiber.Ctx) error {
	tasks, err := c.taskRepo.GetAllTasks()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(tasks)
}

func (c *tasksController) UpdateTask(ctx *fiber.Ctx) error {
	taskID := ctx.Params("id")
	var update models.Task
	err := ctx.BodyParser(&update)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	update.UpdatedAt = time.Now()
	update.Id = bson.ObjectIdHex(taskID)

	err = c.taskRepo.UpdateTask(&update)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}

	return ctx.
		Status(http.StatusOK).
		JSON(update)
}

func (c *tasksController) DeleteTask(ctx *fiber.Ctx) error {
	taskID := ctx.Params("id")
	if !bson.IsObjectIdHex(taskID) {
		return ctx.Status(http.StatusBadRequest).
			JSON(util.NewJError(errors.New("invalid content id")))
	}

	err := c.taskRepo.DeleteTask(taskID)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	ctx.Set("Entity", taskID)
	return ctx.SendStatus(http.StatusOK)
}
