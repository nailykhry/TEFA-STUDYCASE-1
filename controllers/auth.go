package controllers

import (
	"TEFA-STUDYCASE-1/models"     //O(1)
	"TEFA-STUDYCASE-1/repository" //O(1)
	"TEFA-STUDYCASE-1/security"   //O(1)
	"TEFA-STUDYCASE-1/util"       //O(1)
	"fmt"                         //O(1)
	"log"                         //O(1)
	"net/http"                    //O(1)
	"strings"                     //O(1)
	"time"                        //O(1)

	"github.com/form3tech-oss/jwt-go"    //O(1)
	"github.com/gofiber/fiber/v2"        //O(1)
	"gopkg.in/asaskevich/govalidator.v9" //O(1)
	"gopkg.in/mgo.v2"                    //O(1)
	"gopkg.in/mgo.v2/bson"               //O(1)
)

type AuthController interface {
	SignUp(ctx *fiber.Ctx) error     //O(1)
	SignIn(ctx *fiber.Ctx) error     //O(1)
	GetUser(ctx *fiber.Ctx) error    //O(1)
	GetUsers(ctx *fiber.Ctx) error   //O(1)
	PutUser(ctx *fiber.Ctx) error    //O(1)
	DeleteUser(ctx *fiber.Ctx) error //O(1)
}

type authController struct {
	usersRepo repository.UsersRepository //O(1)
}

func NewAuthController(usersRepo repository.UsersRepository) AuthController {
	return &authController{usersRepo} //O(1)
}

func (c *authController) SignUp(ctx *fiber.Ctx) error {
	var newUser models.User         //O(1)
	err := ctx.BodyParser(&newUser) //O(1)
	if err != nil {
		return ctx. //O(1)
				Status(http.StatusUnprocessableEntity).
				JSON(util.NewJError(err))
	}
	newUser.Email = util.NormalizeEmail(newUser.Email) //O(1)
	if !govalidator.IsEmail(newUser.Email) {           //O(1)
		return ctx. //O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrInvalidEmail))
	}
	exists, err := c.usersRepo.GetByEmail(newUser.Email) //O(1)
	if err == mgo.ErrNotFound {                          //O(1)
		if strings.TrimSpace(newUser.Password) == "" { //O(1)
			return ctx. //O(1)
					Status(http.StatusBadRequest).
					JSON(util.NewJError(util.ErrEmptyPassword))
		}
		newUser.Password, err = security.EncryptPassword(newUser.Password) //O(n)
		if err != nil {                                                    //O(1)
			return ctx. //O(1)
					Status(http.StatusBadRequest).
					JSON(util.NewJError(err))
		}
		newUser.CreatedAt = time.Now()        //O(1)
		newUser.UpdatedAt = newUser.CreatedAt //O(1)
		newUser.Role = "student"              //O(1)
		newUser.Id = bson.NewObjectId()       //O(1)
		err = c.usersRepo.Save(&newUser)      //O(1)
		if err != nil {                       //O(1)
			return ctx. //O(1)
					Status(http.StatusBadRequest).
					JSON(util.NewJError(err))
		}
		return ctx. //O(1)
				Status(http.StatusCreated).
				JSON(newUser)
	}

	if exists != nil { //O(1)
		err = util.ErrEmailAlreadyExists //O(1)
	}

	return ctx. //O(1)
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
}

func (c *authController) SignIn(ctx *fiber.Ctx) error {
	var input models.User         //O(1)
	err := ctx.BodyParser(&input) //O(1)
	if err != nil {               //O(1)
		return ctx. //O(1)
				Status(http.StatusUnprocessableEntity).
				JSON(util.NewJError(err))
	}
	input.Email = util.NormalizeEmail(input.Email)   //O(1)
	user, err := c.usersRepo.GetByEmail(input.Email) //O(1)
	if err != nil {                                  //O(1)
		log.Printf("%s signin failed: %v\n", input.Email, err.Error()) //O(1)
		return ctx.                                                    //O(1)
										Status(http.StatusUnauthorized).
										JSON(util.NewJError(util.ErrInvalidCredentials))
	}
	err = security.VerifyPassword(user.Password, input.Password) //O(n)
	if err != nil {                                              //O(1)
		log.Printf("%s signin failed: %v\n", input.Email, err.Error()) //O(1)
		return ctx.                                                    //O(1)
										Status(http.StatusUnauthorized).
										JSON(util.NewJError(util.ErrInvalidCredentials))
	}
	token, err := security.NewToken(user.Id.Hex(), user.Role) //O(n)
	if err != nil {                                           //O(1)
		log.Printf("%s signin failed: %v\n", input.Email, err.Error()) //O(1)
		return ctx.                                                    //O(1)
										Status(http.StatusUnauthorized).
										JSON(util.NewJError(err))
	}
	return ctx. //O(1)
			Status(http.StatusOK). //O(1)
			JSON(fiber.Map{
			"user":  user,
			"token": fmt.Sprintf("Bearer %s", token),
		})
}

func (c *authController) GetUser(ctx *fiber.Ctx) error {
	payload, err := AuthRequestWithId(ctx) //O(n)
	if err != nil {                        //O(1)
		return ctx. //O(1)
				Status(http.StatusUnauthorized).
				JSON(util.NewJError(err))
	}
	user, err := c.usersRepo.GetById(payload.Id) //O(1)
	if err != nil {                              //O(1)
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx. //O(1)
			Status(http.StatusOK).
			JSON(user)
}

func (c *authController) GetUsers(ctx *fiber.Ctx) error {
	users, err := c.usersRepo.GetAll() //O(n)
	if err != nil {                    //O(1)
		return ctx. //O(1)
				Status(http.StatusInternalServerError).
				JSON(util.NewJError(err))
	}
	return ctx. //O(1)
			Status(http.StatusOK).
			JSON(users)
}

func (c *authController) PutUser(ctx *fiber.Ctx) error {
	payload, err := AuthRequestWithId(ctx) //O(n)
	if err != nil {                        //O(1)
		return ctx. //O(1)
				Status(http.StatusUnauthorized).
				JSON(util.NewJError(err))
	}
	var update models.User        //O(1)
	err = ctx.BodyParser(&update) //O(1)
	if err != nil {               //O(1)
		return ctx. //O(1)
				Status(http.StatusUnprocessableEntity).
				JSON(util.NewJError(err))
	}

	update.UpdatedAt = time.Now()            //O(1)
	update.Id = bson.ObjectIdHex(payload.Id) //O(1)
	err = c.usersRepo.Update(&update)        //O(1)

	if err != nil { //O(1)
		return ctx. //O(1)
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
	}

	return ctx. //O(1)
			Status(http.StatusCreated).
			JSON(update)
}

func (c *authController) DeleteUser(ctx *fiber.Ctx) error {
	payload, err := AuthRequestWithId(ctx) //O(n)
	if err != nil {                        //O(1)
		return ctx. //O(1)
				Status(http.StatusUnauthorized).
				JSON(util.NewJError(err))
	}
	err = c.usersRepo.Delete(payload.Id) //O(1)
	if err != nil {                      //O(1)
		return ctx. //O(1)
				Status(http.StatusInternalServerError).
				JSON(util.NewJError(err))
	}
	ctx.Set("Entity", payload.Id)               //O(1)
	return ctx.SendStatus(http.StatusNoContent) //O(1)
}

func AuthRequestWithId(ctx *fiber.Ctx) (*jwt.StandardClaims, error) {
	id := ctx.Params("id")       //O(1)
	if !bson.IsObjectIdHex(id) { //O(1)
		return nil, util.ErrUnauthorized //O(1)
	}
	token := ctx.Locals("user").(*jwt.Token)       //O(1)
	payload, err := security.ParseToken(token.Raw) //O(n)
	if err != nil {                                //O(1)
		return nil, err //O(1)
	}
	if payload.Id != id || payload.Issuer != id { //O(1)
		return nil, util.ErrUnauthorized //O(1)
	}
	return payload, nil //O(1)
}

func PayloadID(ctx *fiber.Ctx) (*jwt.StandardClaims, error) {
	token := ctx.Locals("user").(*jwt.Token)       //O(1)
	payload, err := security.ParseToken(token.Raw) //O(n)
	if err != nil {                                //O(1)
		return nil, err //O(1)
	}
	return payload, nil //O(1)
}
