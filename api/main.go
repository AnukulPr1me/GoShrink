package main

import(
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/anukulpr1me/GoShrink/routes"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.app){
	app.get("/:url", routes.ResolveURL)
	app.post("/api/v1", routes.ShortenURL)
}

func main(){
	err:=godotenv.Load()
	if err!=nil {
        fmt.println(err)
    }
	app := fiber.New()
	app.use(logger.New())
	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}

