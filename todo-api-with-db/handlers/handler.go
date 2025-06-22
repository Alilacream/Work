package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"Todo/db"
	"strconv"
)
// Create CRUD : uses the database handler InsertTodo, which INSERTS the title givin for the todo
func CreateTodoHandler(c *fiber.Ctx) error {
	var todo db.Todo
	
	if err := c.BodyParser(&todo) ;err != nil {
			return c.Status(400).JSON(fiber.Map{"Error": "Invalid Input"})
		}
		insrt := db.InsertTodo(todo.Title) 
		if insrt != nil {
			return c.Status(500).JSON(fiber.Map{"Error" : "Couldn't Insert Todo"})
		}
		return c.Status(200).JSON(todo)
}
// Delete CRUD : uses the database handler DeleteTodo, which checks where the id so it can delete everything from that id 
func DeleteTodoHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	id_INT, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error": "Invalid ID"})
	}
	err = db.DeleteTodo(id_INT)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error": "Couldn't Delete the todo"})
	}
	return c.Status(200).JSON(fiber.Map{"Message": fmt.Sprintf("Todo of id = %d is succesfully deleted ", id_INT)})
}
// Read-one CRUD : uses the GetTodo db hand, uses the id to check for the todo to read it 
func GetTodoHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	id_INT , err := strconv.Atoi(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error": "Invalid ID"})
	}
	todo, err := db.GetTodo(id_INT)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error":"Couldn't get the todo"})
	}
	return c.Status(200).JSON(todo)
}
// Read-All CRUD : uses the GetAllTodos, returns a slice of todos and an error, now we just show them by using the range of todos
func GetAllTodosHandler(c *fiber.Ctx) error {
	todos, err := db.GetAllTodos()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error": "Unable to get all todos"})
	}	
	for i, todo := range todos {
		if (todo.Title == "") {
			return c.Status(400).JSON(fiber.Map{"Missing Title": fmt.Sprintf("The todo of id: %d is missing it's purpose (title)", i)})
		}
	}
		return c.Status(200).JSON(todos)
}
// Update CRUD : updates the Title and Done attribute of the givin Id
func UpdateTodoHandler(c *fiber.Ctx) error {
	var todo db.Todo

	id := c.Params("id")	
	id_INT ,err:=  strconv.Atoi(id)
	
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error": "Invalid ID"})
	}
	if err:= c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"Error":"Invalid Input"})
	}
	err = db.UpdateTodo(id_INT, todo.Title, todo.Done)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error":"Failed to update todo"})
	}
	return c.Status(201).JSON(todo)
}

