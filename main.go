package main

import (
	"TEFA-STUDYCASE-1/controllers"
	"TEFA-STUDYCASE-1/database"
	"TEFA-STUDYCASE-1/repository"
	"TEFA-STUDYCASE-1/routes"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	conn := database.NewConnection()
	defer conn.Close()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Hello World"})
	})

	usersRepo := repository.NewUsersRepository(conn)
	authController := controllers.NewAuthController(usersRepo)
	authRoutes := routes.NewAuthRoutes(authController)
	authRoutes.Install(app)

	contentsRepo := repository.NewContentRepository(conn)
	contentsController := controllers.NewContentController(contentsRepo)
	contentRoutes := routes.NewContentRoutes(contentsController)
	contentRoutes.Content(app)

	tasksRepo := repository.NewTaskRepository(conn)
	tasksController := controllers.NewTaskController(tasksRepo)
	tasksRoutes := routes.NewTaskRoutes(tasksController)
	tasksRoutes.Task(app)

	usersubRepo := repository.NewUsersubRepository(conn)
	usersubController := controllers.NewUsersubController(usersubRepo)
	usersubRoutes := routes.NewUsersubRoutes(usersubController)
	usersubRoutes.Usersub(app)

	userchatRepo := repository.NewTUserchatRepository(conn)
	userchatController := controllers.NewUserchatController(userchatRepo)
	userchatRoutes := routes.NewUserchatRoutes(userchatController)
	userchatRoutes.Userchat(app)

	log.Fatal(app.Listen(":8080"))
}
