package main

import (
	"log"

	"backend-task/config"
	"backend-task/internal/handler"
	"backend-task/internal/repository"
	"backend-task/internal/routes"
	"backend-task/internal/service"

	db "backend-task/db/sqlc"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	dbConn := config.ConnectDatabase()
	defer dbConn.Close()

	queries := db.New(dbConn)

	repo := repository.NewUserRepository(queries)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	routes.Register(app, handler)

	log.Fatal(app.Listen(":3000"))
}
