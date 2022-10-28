package routes

import (
	controllers "go-mongo-api2/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	//USER GROUP
	user := api.Group("/user")
	user.Post("/", controllers.UserStore)
	user.Get("/gets", controllers.UserIndex)
	user.Get("/:id", controllers.UserGetId)
	user.Post("update/:id", controllers.UserUpdate)
	user.Delete("/:id", controllers.UserDestroy)

	//USER GROUP END

}
