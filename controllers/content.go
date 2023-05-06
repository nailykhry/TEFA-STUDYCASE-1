package controllers

import (
	"TEFA-STUDYCASE-1/models"     //O(1)
	"TEFA-STUDYCASE-1/repository" //O(1)
	"TEFA-STUDYCASE-1/security"
	"TEFA-STUDYCASE-1/util" //O(1)
	"errors"                //O(1)
	"net/http"              //O(1)
	"time"                  //O(1)

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2" //O(1)
	"gopkg.in/mgo.v2/bson"        //O(1)
)

type ContentsController interface {
	UploadContent(ctx *fiber.Ctx) error //O(1)
	GetContent(ctx *fiber.Ctx) error    //O(1)
	GetContents(ctx *fiber.Ctx) error   //O(1)
	PutContent(ctx *fiber.Ctx) error    //O(1)
	DeleteContent(ctx *fiber.Ctx) error //O(1)
}

func NewContentController(contentsRepo repository.ContentsRepository) ContentsController {
	return &contentsController{contentsRepo} //O(1)
}

type contentsController struct {
	contentRepo repository.ContentsRepository //O(1)
}

func (c *contentsController) UploadContent(ctx *fiber.Ctx) error {
	payload, err := PayloadID(ctx)
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(err))
	}
	var newContent models.Content
	err = ctx.BodyParser(&newContent)
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
	newContent.User_ID = payload.Id
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
	contentID := ctx.Params("id")
	if !bson.IsObjectIdHex(contentID) {
		return ctx.Status(http.StatusBadRequest).
			JSON(util.NewJError(errors.New("invalid content id")))
	}

	content, err := c.contentRepo.GetById(contentID)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	//cek authorize
	err = ContentRequestWithId(ctx, content.User_ID)
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(content)
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

	//cek authorize
	err = ContentRequestWithId(ctx, update.User_ID)
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(err))
	}

	update.UpdatedAt = time.Now()
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

// AUTHORIZATION
func ContentRequestWithId(ctx *fiber.Ctx, id string) error {
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		return err
	}

	if payload.Id != id || payload.Issuer != id {
		return util.ErrUnauthorized
	}
	return nil
}
