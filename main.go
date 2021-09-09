package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Person struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	users := []*Person{
		{
			Name: "John",
			Id:   1,
		},
		{
			Name: "Yashira",
			Id:   2,
		},
		{
			Name: "Timo",
			Id:   3,
		},
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"results": users,
		})
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			fmt.Println("There was an error! :(")
		}

		user := users[id-1]

		return c.JSON(user)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		type Request struct {
			Name string `json:"name"`
		}

		var body Request

		// BodyParser is a method that injects the content of the body in a variable, its first parameter is the reference of the variable in which you wanna save the content of the body, but previvously you must declare an struct with the keys, and asign it to the variable
		err := c.BodyParser(&body)

		// if error
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Cannot parse JSON",
			})
		}

		// create a newName variable
		newName := &Person{
			Name: body.Name,
			Id:   len(users) + 1,
		}

		// append in users
		users = append(users, newName)

		// fmt.Println(users)
		return c.JSON(fiber.Map{
			"results": newName,
		})
	})

	app.Put("/:id", func(c *fiber.Ctx) error {
		paramsId, err := c.ParamsInt("id")

		if err != nil {
			fmt.Println("There was an error!")
		}

		type Request struct {
			Name *string `json:"name"`
		}

		var body Request

		c.BodyParser(&body)

		var user *Person

		for _, u := range users {
			if u.Id == paramsId {
				user = u
				break
			}
		}

		fmt.Println(*body.Name)

		user.Name = *body.Name

		return c.JSON(fiber.Map{
			"results": users,
		})
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {

		paramsId, _ := c.ParamsInt("id")

		for i, user := range users {
			if user.Id == paramsId {
				users = append(users[0:i], users[i+1:]...)

				return c.JSON(fiber.Map{
					"success": true,
					"message": "Deleted Successfully",
				})
			}
		}
		// if todo not found
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "user not found",
		})

	})

	app.Listen(":3000")
}
