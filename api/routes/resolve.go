package routes 

import (
	"github.com/anukulpr1me/GoShrink/database"
	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/joho/godotenv"
)

func ResolveURL(c *fiber.Ctx) error{
	url:= c.Params("url")
	r:=database.CreateClient(0)
	defer r.Close()
	value, err := r.Get(database.Ctx, url).Result()
	if err == radis.Nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "URL not found",
        })
	}else if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_=rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)
}