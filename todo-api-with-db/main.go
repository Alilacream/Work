package main

import (
	"Todo/db"
	"Todo/handlers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()	
	db.Connect()
	PORT := os.Getenv("PORT")
	app := fiber.New()
	app.Post("/todo", handlers.CreateTodoHandler)
	app.Get("/todo/all", handlers.GetAllTodosHandler)
	app.Get("/todo/:id" , handlers.GetTodoHandler)
	app.Patch("/todo/:id" , handlers.UpdateTodoHandler)
	app.Delete("/todo/:id", handlers.DeleteTodoHandler)
	log.Fatal(app.Listen(":"+PORT))
}