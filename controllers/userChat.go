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

type UserchatsController interface {
	CreateUserchat(ctx *fiber.Ctx) error // O(1)
	GetUserchat(ctx *fiber.Ctx) error    // O(1)
	GetUserchats(ctx *fiber.Ctx) error   // O(1)
	UpdateUserchat(ctx *fiber.Ctx) error // O(1)
}

func NewUserchatController(userchatsRepo repository.UserchatsRepository) UserchatsController {
	return &userchatsController{userchatsRepo} // O(1)
}

type userchatsController struct {
	userchatRepo repository.UserchatsRepository // O(1)
}

func (c *userchatsController) CreateUserchat(ctx *fiber.Ctx) error {
	var newUserchat models.Userchat     // O(1)
	err := ctx.BodyParser(&newUserchat) // O(1)
	if err != nil {                     // O(1)
		return ctx. // O(1)
				Status(http.StatusUnprocessableEntity).
				JSON(util.NewJError(err))
	}

	if newUserchat.Title == "" || newUserchat.Question == "" { // O(1)
		return ctx. // O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(errors.New("bad request: invalid userchat")))
	}

	newUserchat.Status = false                    // O(1)
	newUserchat.CreatedAt = time.Now()            // O(1)
	newUserchat.UpdatedAt = newUserchat.CreatedAt // O(1)
	newUserchat.Id = bson.NewObjectId()           // O(1)

	err = c.userchatRepo.CreateUserchat(&newUserchat) // O(1)
	if err != nil {                                   // O(1)
		return ctx. // O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
	}
	return ctx. // O(1)
			Status(http.StatusCreated).
			JSON(newUserchat)
}

func (c *userchatsController) GetUserchat(ctx *fiber.Ctx) error {
	userchatID := ctx.Params("id")       // O(1)
	if !bson.IsObjectIdHex(userchatID) { // O(1)
		return ctx.Status(http.StatusBadRequest). // O(1)
								JSON(util.NewJError(errors.New("invalid userchat id")))
	}
	user, err := c.userchatRepo.GetUserchatById(userchatID) // O(1)
	if err != nil {                                         // O(1)
		return ctx. // O(1)
				Status(http.StatusInternalServerError).
				JSON(util.NewJError(err))
	}
	return ctx. // O(1)
			Status(http.StatusOK).
			JSON(user)
}

func (c *userchatsController) GetUserchats(ctx *fiber.Ctx) error {
	userchats, err := c.userchatRepo.GetAllUserchats() // O(n)
	if err != nil {                                    // O(1)
		return ctx. // O(1)
				Status(http.StatusInternalServerError).
				JSON(util.NewJError(err))
	}
	return ctx. // O(1)
			Status(http.StatusOK).
			JSON(userchats)
}

func (c *userchatsController) UpdateUserchat(ctx *fiber.Ctx) error {
	userchatID := ctx.Params("id") // O(1)
	var update models.Userchat     // O(1)
	err := ctx.BodyParser(&update) // O(1)
	if err != nil {                // O(1)
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
	update.Status = true                     // O(1)
	update.UpdatedAt = time.Now()            // O(1)
	update.Id = bson.ObjectIdHex(userchatID) // O(1)

	err = c.userchatRepo.UpdateUserchat(&update) // O(1)
	if err != nil {                              // O(1)
		return ctx. // O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
	}

	return ctx. // O(1)
			Status(http.StatusOK).
			JSON(update)
}
