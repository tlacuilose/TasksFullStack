package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/tlacuilose/tasks-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Get all todos - GET /api/todos
func GetAllTodos(ctx *fiber.Ctx) error {
	// Get mongodb collection.
	collection := mgm.Coll(&models.Todo{})

	// Array to store all todos.
	todos := []models.Todo{}

	// Retrieve all todos from db.
	err := collection.SimpleFind(&todos, bson.D{})
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return nil // stop execution if error.
	}

	// Return results.
	ctx.JSON(fiber.Map{
		"ok":    true,
		"todos": todos,
	})
	return nil
}

// Get a todo by id - GET /api/todos/:id
func GetTodoByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return nil
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
	return nil
}

// Create a new todo - POST /api/todos
func CreateTodo(ctx *fiber.Ctx) error {
	// Define a new structure to comply request info
	params := new(struct {
		Title       string
		Description string
	})

	ctx.BodyParser(&params) // Check if the body request complies with stuct.

	if len(params.Title) == 0 || len(params.Description) == 0 {
		ctx.Status(400).JSON(fiber.Map{ // If it doesnt return an error.
			"ok":    false,
			"error": "Title or description not specified.",
		})
		return nil
	}

	todo := models.CreateTodo(params.Title, params.Description)
	err := mgm.Coll(todo).Create(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return nil
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
	return nil
}

// Update a todo done status - PATCH /api/todos/:id
func ToggleTodoStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return nil
	}

	todo.Done = !todo.Done

	err = collection.Update(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return nil
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
	return nil
}

// Delete todo - DELETE /api/todos/:id
func DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return nil
	}

	err = collection.Delete(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return nil
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
	return nil
}
