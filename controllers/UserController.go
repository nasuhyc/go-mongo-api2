package controllers

import (
	"context"
	"go-mongo-api2/config"
	"go-mongo-api2/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserStore(c *fiber.Ctx) error {

	userCollection := config.MI.DB.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(models.User)
	user.Id = primitive.NewObjectID()

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error:":  err,
		})
	}
	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "User failed to insert",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "User inserted successfully",
	})

}

func UserIndex(c *fiber.Ctx) error {
	var user models.User
	var users []models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := config.MI.DB.Collection("user").Find(ctx, bson.M{})

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	for result.Next(ctx) {
		if err := result.Decode(&user); err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    users,
		"message": "Registered users listed",
	})

}

func UserGetId(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User

	dataId, err := primitive.ObjectIDFromHex(c.Params("id"))
	getData := userCollection.FindOne(ctx, bson.M{"_id": dataId})
	if err := getData.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Users Not found",
			"error":   err,
		})
	}
	err = getData.Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Users Not found",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"success": true,
	})

}

func UserUpdate(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "failed to parse body",
			"error":   err,
		})
	}
	getId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not Found",
			"error":   err,
		})
	}

	update := bson.M{
		"$set": user,
	}
	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": getId}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User updated successfully",
		"data":    user,
	})

}

func UserDestroy(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	getId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"error":   err,
		})
	}
	_, err = userCollection.DeleteOne(ctx, bson.M{"_id": getId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "User failed to delete",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully",
	})

}
