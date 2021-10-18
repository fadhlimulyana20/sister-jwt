package router

import (
	"github.com/fadhlimulyana20/sister-jwt/controller"
	m "github.com/fadhlimulyana20/sister-jwt/middleware"
	"github.com/gofiber/fiber/v2"
)

type Api struct{}

func (a *Api) Init(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Get("/me", controller.Me)

	pbb := app.Group("/pbb", m.JwtMiddleware)
	pbb.Get("/", controller.GetPbb)
	pbb.Post("/create", controller.CreatePbb)
	pbb.Put("/add_data_pbb/:id", controller.AddDataPbb)
}
