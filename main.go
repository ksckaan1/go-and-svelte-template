package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	//template render engine
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine, //set as render engine
	})
	app.Static("/public", "./public")
	app.Get("/", mainPage)
	app.Listen(":3000")
}

func mainPage(c *fiber.Ctx) error {
	return c.Render("mainpage", fiber.Map{})
}
