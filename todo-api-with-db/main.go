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
	db.Connect()
	godotenv.Load()	
	PORT := os.Getenv("PORT")
	app := fiber.New()
	defer log.Fatal(app.Listen(":"+PORT))
	app.Get("/todo", handlers.CreateTodoHandler)
	app.Get("/todo/all", handlers.GetAllTodosHandler)
	app.Get("/todo/:id" , handlers.GetTodoHandler)
	app.Patch("/todo/:id" , handlers.UpdateTodoHandler)
	app.Delete("/todo/:id", handlers.DeleteTodoHandler)
}