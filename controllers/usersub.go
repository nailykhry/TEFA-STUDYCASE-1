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

type UsersubsController interface {
	CreateUsersub(ctx *fiber.Ctx) error
}

func NewUsersubController(usersubsRepo repository.UsersubsRepository) UsersubsController {
	return &usersubsController{usersubsRepo}
}

type usersubsController struct {
	usersubRepo repository.UsersubsRepository
}

func (c *usersubsController) CreateUsersub(ctx *fiber.Ctx) error {
	var newUsersub models.Usersub
	err := ctx.BodyParser(&newUsersub)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	if newUsersub.Plan == "" || newUsersub.Price == 0 {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(errors.New("bad request: invalid task")))
	}

	newUsersub.CreatedAt = time.Now()
	newUsersub.Id = bson.NewObjectId()

	err = c.usersubRepo.CreateUsersub(&newUsersub)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusCreated).
		JSON(newUsersub)
}
