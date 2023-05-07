package controllers

import (
	"TEFA-STUDYCASE-1/models"     //O(1)
	"TEFA-STUDYCASE-1/repository" //O(1)
	"TEFA-STUDYCASE-1/util"       //O(1)
	"errors"                      //O(1)
	"net/http"                    //O(1)
	"time"                        //O(1)

	"github.com/gofiber/fiber/v2" //O(1)
	"gopkg.in/mgo.v2/bson"        //O(1)
)

type UsersubsController interface {
	CreateUsersub(ctx *fiber.Ctx) error //O(1)
}

func NewUsersubController(usersubsRepo repository.UsersubsRepository) UsersubsController {
	return &usersubsController{usersubsRepo} //O(1)
}

type usersubsController struct {
	usersubRepo repository.UsersubsRepository //O(1)
}

func (c *usersubsController) CreateUsersub(ctx *fiber.Ctx) error {
	var newUsersub models.Usersub      //O(1)
	err := ctx.BodyParser(&newUsersub) //O(1)
	if err != nil {                    //O(1)
		return ctx. //O(1)
				Status(http.StatusUnprocessableEntity).
				JSON(util.NewJError(err))
	}

	if newUsersub.Plan == "" || newUsersub.Price == 0 { //O(1)
		return ctx. //O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(errors.New("bad request: invalid task")))
	}

	newUsersub.CreatedAt = time.Now()  //O(1)
	newUsersub.Id = bson.NewObjectId() //O(1)

	err = c.usersubRepo.CreateUsersub(&newUsersub) //O(1)
	if err != nil {                                //O(1)
		return ctx. //O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
	}
	return ctx. //O(1)
			Status(http.StatusCreated).
			JSON(newUsersub)
}
