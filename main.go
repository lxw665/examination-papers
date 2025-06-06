package main

import (
	"examination-papers/configs"
	"examination-papers/controllers"
	"examination-papers/data/db"
	"examination-papers/data/redis"
	"examination-papers/middleware"
	"examination-papers/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//app := fiber.New()
	//
	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, World!")
	//})
	//
	//app.Listen(":3000")
	dbConfig := db.Config{
		Host:     "localhost",
		Port:     25432,
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	}
	config := configs.FiberConfig()
	app := fiber.New(config)
	middleware.FiberMiddleware(app)
	dbClient, err := db.NewPostgresClient(dbConfig)
	if err != nil {
		panic(err)
	}
	redisConfig := redis.Config{
		Host:     "localhost",
		Port:     6379,
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       0,
	}
	redisClient, err := redis.NewRedisClient(redisConfig)
	if err != nil {
		panic(err)
	}
	examCase := controllers.NewSubmitExamCase(dbClient.DB, nil, redisClient.Client)
	go examCase.SubmitExamWorker()
	routes.PublicRoutes(app, examCase)

	app.Listen(":3001")
}
