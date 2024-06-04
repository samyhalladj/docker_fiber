package main

import (
	"os/exec"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("en fonction")
	})

	app.Get("/pi", func(c fiber.Ctx) error {
		// Exécuter le script shell
		out, err := exec.Command("/bin/sh", "./pi.sh").Output()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		// Retourner le résultat du script shell
		return c.SendString(string(out))
	})

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
