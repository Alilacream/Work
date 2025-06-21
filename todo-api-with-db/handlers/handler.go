package handlers

import (
	"github.com/gofiber/fiber/v2"
	"Todo/db"
	"strconv"
)

func CreateTodoHandler(c *fiber.Ctx) error {
	var todo db.Todo
	
	if err := c.BodyParser(&todo) ;err != nil {
			return c.Status(400).JSON(fiber.Map{"Error": fmt.Sprint(err)})
		}
		insrt := db.InsertTodo(todo.Title) 
		if insrt != nil {
			return c.Status(500).JSON(fiber.Map{"Error" : fmt.Sprint(insrt)})
		}
		return c.Status(200).JSON(todo)
}

func UpdateTodoHandler(c *fiber.Ctx) error {

}

