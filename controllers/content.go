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

type ContentsController interface {
	UploadContent(ctx *fiber.Ctx) error
	GetContent(ctx *fiber.Ctx) error
	GetContents(ctx *fiber.Ctx) error
	PutContent(ctx *fiber.Ctx) error
	DeleteContent(ctx *fiber.Ctx) error
}

func NewContentController(contentsRepo repository.ContentsRepository) ContentsController {
	return &contentsController{contentsRepo}
}

type contentsController struct {
	contentRepo repository.ContentsRepository
}

func (c *contentsController) UploadContent(ctx *fiber.Ctx) error {
	var newContent models.Content
	err := ctx.BodyParser(&newContent)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	if newContent.Title == "" || newContent.Content == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(errors.New("bad request: invalid content")))
	}

	newContent.CreatedAt = time.Now()
	newContent.UpdatedAt = newContent.CreatedAt
	newContent.User_ID = "123456"
	newContent.Id = bson.NewObjectId()
	err = c.contentRepo.UploadContent(&newContent)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusCreated).
		JSON(newContent)
}

func (c *contentsController) GetContent(ctx *fiber.Ctx) error {
	// payload, err := AuthRequestWithId(ctx)
	// if err != nil {
	// 	return ctx.
	// 		Status(http.StatusUnauthorized).
	// 		JSON(util.NewJError(err))
	// }

	contentID := ctx.Params("id")
	if !bson.IsObjectIdHex(contentID) {
		return ctx.Status(http.StatusBadRequest).
			JSON(util.NewJError(errors.New("invalid content id")))
	}

	user, err := c.contentRepo.GetById(contentID)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(user)
}

func (c *contentsController) GetContents(ctx *fiber.Ctx) error {
	contents, err := c.contentRepo.GetAll()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(contents)
}

func (c *contentsController) PutContent(ctx *fiber.Ctx) error {
	// payload, err := AuthRequestWithId(ctx)
	// if err != nil {
	// 	return ctx.
	// 		Status(http.StatusUnauthorized).
	// 		JSON(util.NewJError(err))
	// }
	contentID := ctx.Params("id")
	if !bson.IsObjectIdHex(contentID) {
		return ctx.Status(http.StatusBadRequest).
			JSON(util.NewJError(errors.New("invalid content id")))
	}
	var update models.Content
	err := ctx.BodyParser(&update)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	update.UpdatedAt = time.Now()
	update.User_ID = "123456"
	update.Id = bson.ObjectIdHex(contentID)
	err = c.contentRepo.Update(&update)
	print(update.Title)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}

	return ctx.
		Status(http.StatusCreated).
		JSON(update)
}

func (c *contentsController) DeleteContent(ctx *fiber.Ctx) error {
	// payload, err := AuthRequestWithId(ctx)
	// if err != nil {
	// 	return ctx.
	// 		Status(http.StatusUnauthorized).
	// 		JSON(util.NewJError(err))
	// }
	contentID := ctx.Params("id")
	if !bson.IsObjectIdHex(contentID) {
		return ctx.Status(http.StatusBadRequest).
			JSON(util.NewJError(errors.New("invalid content id")))
	}

	err := c.contentRepo.Delete(contentID)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	ctx.Set("Entity", contentID)
	return ctx.SendStatus(http.StatusOK)
}
