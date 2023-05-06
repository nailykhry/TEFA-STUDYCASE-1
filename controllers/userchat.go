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

type UserchatsController interface {
	CreateUserchat(ctx *fiber.Ctx) error
	GetUserchat(ctx *fiber.Ctx) error
	GetUserchats(ctx *fiber.Ctx) error
	UpdateUserchat(ctx *fiber.Ctx) error
}

func NewUserchatController(userchatsRepo repository.UserchatsRepository) UserchatsController {
	return &userchatsController{userchatsRepo}
}

type userchatsController struct {
	userchatRepo repository.UserchatsRepository
}

func (c *userchatsController) CreateUserchat(ctx *fiber.Ctx) error {
	var newUserchat models.Userchat
	err := ctx.BodyParser(&newUserchat)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	if newUserchat.Title == "" || newUserchat.Question == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(errors.New("bad request: invalid chat")))
	}

	newUserchat.Status = false
	newUserchat.CreatedAt = time.Now()
	newUserchat.UpdatedAt = newUserchat.CreatedAt
	newUserchat.Id = bson.NewObjectId()

	err = c.userchatRepo.CreateUserchat(&newUserchat)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusCreated).
		JSON(newUserchat)
}

func (c *userchatsController) GetUserchat(ctx *fiber.Ctx) error {
	userchatID := ctx.Params("id")
	if !bson.IsObjectIdHex(userchatID) {
		return ctx.Status(http.StatusBadRequest).
			JSON(util.NewJError(errors.New("invalid content id")))
	}
	user, err := c.userchatRepo.GetUserchatById(userchatID)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(user)
}

func (c *userchatsController) GetUserchats(ctx *fiber.Ctx) error {
	userchats, err := c.userchatRepo.GetAllUserchats()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(userchats)
}

func (c *userchatsController) UpdateUserchat(ctx *fiber.Ctx) error {
	userchatID := ctx.Params("id")
	var update models.Userchat
	err := ctx.BodyParser(&update)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
	update.Status = true
	update.UpdatedAt = time.Now()
	update.Id = bson.ObjectIdHex(userchatID)

	err = c.userchatRepo.UpdateUserchat(&update)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}

	return ctx.
		Status(http.StatusOK).
		JSON(update)
}
