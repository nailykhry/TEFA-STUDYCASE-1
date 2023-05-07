package controllers

import (
	"TEFA-STUDYCASE-1/models"     //O(1)
	"TEFA-STUDYCASE-1/repository" //O(1)
	"TEFA-STUDYCASE-1/security"
	"TEFA-STUDYCASE-1/util" //O(1)
	"errors"                //O(1)
	"net/http"              //O(1)
	"time"                  //O(1)

	"github.com/form3tech-oss/jwt-go" //O(1)
	"github.com/gofiber/fiber/v2"     //O(1)
	"gopkg.in/mgo.v2/bson"            //O(1)
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
	payload, err := PayloadID(ctx) //O(n)
	if err != nil {                //O(1)
		return ctx. //O(1)
				Status(http.StatusUnauthorized).
				JSON(util.NewJError(err))
	}
	var newContent models.Content     //O(1)
	err = ctx.BodyParser(&newContent) //O(1)
	if err != nil {                   //O(1)
		return ctx. //O(1)
				Status(http.StatusUnprocessableEntity).
				JSON(util.NewJError(err))
	}

	if newContent.Title == "" || newContent.Content == "" { //O(1)
		return ctx. //O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(errors.New("bad request: invalid content")))
	}

	newContent.CreatedAt = time.Now()              //O(1)
	newContent.UpdatedAt = newContent.CreatedAt    //O(1)
	newContent.User_ID = payload.Id                //O(1)
	newContent.Id = bson.NewObjectId()             //O(1)
	err = c.contentRepo.UploadContent(&newContent) //O(1)
	if err != nil {                                //O(1)
		return ctx. //O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
	}
	return ctx. //O(1)
			Status(http.StatusCreated).
			JSON(newContent)
}

func (c *contentsController) GetContent(ctx *fiber.Ctx) error {
	contentID := ctx.Params("id")       //O(1)
	if !bson.IsObjectIdHex(contentID) { //O(1)
		return ctx.Status(http.StatusBadRequest). //O(1)
								JSON(util.NewJError(errors.New("invalid content id")))
	}

	content, err := c.contentRepo.GetById(contentID) //O(1)
	if err != nil {                                  //O(1)
		return ctx. //O(1)
				Status(http.StatusInternalServerError).
				JSON(util.NewJError(err))
	}
	//cek authorize
	err = ContentRequestWithId(ctx, content.User_ID) //O(n)
	if err != nil {                                  //O(1)
		return ctx. //O(1)
				Status(http.StatusUnauthorized).
				JSON(util.NewJError(err))
	}
	return ctx. //O(1)
			Status(http.StatusOK). //O(1)
			JSON(content)
}

func (c *contentsController) GetContents(ctx *fiber.Ctx) error {
	contents, err := c.contentRepo.GetAll() //O(n)
	if err != nil {                         //O(1)
		return ctx. //O(1)
				Status(http.StatusInternalServerError).
				JSON(util.NewJError(err))
	}
	return ctx. //O(1)
			Status(http.StatusOK).
			JSON(contents)
}

func (c *contentsController) PutContent(ctx *fiber.Ctx) error {
	contentID := ctx.Params("id")       //O(1)
	if !bson.IsObjectIdHex(contentID) { //O(1)
		return ctx.Status(http.StatusBadRequest). //O(1)
								JSON(util.NewJError(errors.New("invalid content id")))
	}
	var update models.Content      //O(1)
	err := ctx.BodyParser(&update) //O(1)
	if err != nil {                //O(1)
		return ctx. //O(1)
				Status(http.StatusUnprocessableEntity).
				JSON(util.NewJError(err))
	}

	//cek authorize
	err = ContentRequestWithId(ctx, update.User_ID) //O(n)
	if err != nil {                                 //O(1)
		return ctx. //O(1)
				Status(http.StatusUnauthorized).
				JSON(util.NewJError(err))
	}

	update.UpdatedAt = time.Now()           //O(1)
	update.Id = bson.ObjectIdHex(contentID) //O(1)
	err = c.contentRepo.Update(&update)     //O(1)

	if err != nil { //O(1)
		return ctx. //O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
	}

	return ctx. //O(1)
			Status(http.StatusCreated).
			JSON(update)
}

func (c *contentsController) DeleteContent(ctx *fiber.Ctx) error {
	contentID := ctx.Params("id")       //O(1)
	if !bson.IsObjectIdHex(contentID) { //O(1)
		return ctx.Status(http.StatusBadRequest). //O(1)
								JSON(util.NewJError(errors.New("invalid content id")))
	}

	err := c.contentRepo.Delete(contentID) //O(1)
	if err != nil {                        //O(1)
		return ctx. //O(1)
				Status(http.StatusInternalServerError).
				JSON(util.NewJError(err))
	}
	ctx.Set("Entity", contentID)         //O(1)
	return ctx.SendStatus(http.StatusOK) //O(1)
}

// AUTHORIZATION
func ContentRequestWithId(ctx *fiber.Ctx, id string) error {
	token := ctx.Locals("user").(*jwt.Token)       //O(1)
	payload, err := security.ParseToken(token.Raw) //O(n)
	if err != nil {                                //O(1)
		return err //O(1)
	}

	if payload.Id != id || payload.Issuer != id { //O(1)
		return util.ErrUnauthorized //O(1)
	}
	return nil //O(1)
}
