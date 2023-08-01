package User

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"rateService/prisma/db"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func DeleteUserById(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")
	client := db.NewClient()
	client.Prisma.Connect()
	result, err := client.User.FindUnique(
		db.User.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(result)
}

func GetUsers(c *fiber.Ctx) error {
	ctx := context.Background()
	client := db.NewClient()
	client.Prisma.Connect()
	result, err := client.User.FindMany().Exec(ctx)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(result)
}

func GetUserById(c *fiber.Ctx) error {
	ctx := context.Background()
	client := db.NewClient()
	client.Prisma.Connect()
	id := c.Params("id")
	result, err := client.User.FindUnique(
		db.User.ID.Equals(id)).Exec(ctx)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(result)
}

func CreateUser(c *fiber.Ctx) error {
	var newUser User
	ctx := context.Background()
	client := db.NewClient()
	client.Prisma.Connect()
	result, err := client.User.CreateOne(
		db.User.Email.Set(newUser.Email),
		db.User.Name.Set(newUser.Name),
		db.User.Password.Set(newUser.Password),
	).Exec(ctx)
	if err != nil {
		return err
	}
	println(result)
	return c.Status(201).JSON(result)
}
func UpdateById(c *fiber.Ctx) error {
	ctx := context.Background()
	client := db.NewClient()
	client.Prisma.Connect()
	id := c.Params("id")
	var updatedUser User
	result, err := client.User.FindUnique(
		db.User.ID.Equals(id)).Update(
		db.User.Email.Set(updatedUser.Email),
		db.User.Name.Set(updatedUser.Name),
		db.User.Password.Set(updatedUser.Password),
	).Exec(ctx)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(result)
}
