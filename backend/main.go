package main

import (
	"math/rand"
	"fmt"
     "time"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	  ID    int    `json:"id"`
	  Name  string `json:"name"`
	  Task string `json:"task"`
	  Completed bool `json:"completed"`
}
type Response struct {
	  Name  string `json:"name"`
	  Message string `json:"message"`
}
 var user []User
func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  app.Post("/add", func(c *fiber.Ctx) error {
    resp:=new(Response)
	if err := c.BodyParser(resp); err != nil {
		return err
	}
	fmt.Println(resp.Name)
	fmt.Println(resp.Message)
	rand.Seed(time.Now().UnixNano())

	// Generate a random 3-digit number (between 100 and 999)
	id := rand.Intn(900) + 100
	user = append(user, User{ID: id, Name: resp.Name, Task: resp.Message, Completed: false})
	return c.JSON(fiber.Map{
		"data": user,
		"message": "Task Added Successfully",
	})
  })

  app.Get("/list", func(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": user,
		"message": "List of Tasks",
	})
	})

	app.Put("/update/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}
		for i := range user{
			if user[i].ID == id {
				user[i].Completed = true
				return c.JSON(fiber.Map{
					"data": user,
					"message": "Task Updated Successfully",
				})
			}
		}
		return c.Status(404).JSON(fiber.Map{
			"error": "Task Not Found",
		})
	});
	app.Delete("/delete/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}
		for i := range user{
			if user[i].ID == id {
				user = append(user[:i], user[i+1:]...)
				return c.JSON(fiber.Map{
					"data": user,
					"message": "Task Deleted Successfully",
				})
			}
		}
		return c.Status(404).JSON(fiber.Map{
			"error": "Task Not Found",
		})
	})

	app.Listen(":3001")
}
