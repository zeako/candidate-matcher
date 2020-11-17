package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/zeako/candidate-matcher/internal/pkg/candidates"
	"github.com/zeako/candidate-matcher/internal/pkg/database"
	"github.com/zeako/candidate-matcher/internal/pkg/jobs"
)

func main() {
	dbpath := os.Getenv("DB_FILE_PATH")
	db, err := database.Init(dbpath)
	if err != nil {
		log.Fatalf("failed opening sqlite db: %v", err)
	}
	defer db.Close()

	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := fiber.New()
	app.Use(logger.New(), recover.New())

	app.Get("/health", func(c *fiber.Ctx) error { return c.SendStatus(200) })

	candidates.Route(app.Group("/candidates"))
	jobs.Route(app.Group("/jobs"))

	log.Fatal(app.Listen(":8080"))
}
