package controller

import (
	"context"
	"net/http"
	"strings"

	"github.com/fadhlimulyana20/sister-jwt/database"
	"github.com/fadhlimulyana20/sister-jwt/helper"
	"github.com/fadhlimulyana20/sister-jwt/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

type loginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	u := new(model.User)

	if err := c.BodyParser(u); err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	u.Id = primitive.NewObjectID()

	db := database.DbManager()
	_, err := db.Collection("users").InsertOne(ctx, u)
	if err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&model.Json{
		Status: "success",
		Data:   u,
	})
}

func Login(c *fiber.Ctx) error {
	l := new(loginDTO)
	u := new(model.User)

	if err := c.BodyParser(l); err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	db := database.DbManager()
	err := db.Collection("users").FindOne(ctx, bson.M{"email": l.Email}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		return c.Status(404).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	} else if err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	if l.Password != u.Password {
		return c.Status(http.StatusUnauthorized).JSON(&model.Json{
			Status:  "failed",
			Message: "Password Salah",
		})
	}

	j := new(helper.JWT)
	token, err := j.CreateToken(u.Id.Hex(), 48, "token")
	if err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&model.Json{
		Status: "success",
		Data: map[string]interface{}{
			"token": token,
			"user":  u,
		},
	})
}

func Me(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	u := new(model.User)

	if !strings.Contains(authHeader, "Bearer") {
		data := &model.Json{
			Status:  "failed",
			Message: "Anda belum Login",
		}

		return c.Status(http.StatusUnauthorized).JSON(data)
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	j := new(helper.JWT)
	userId, err := j.Parse(tokenString, "token")

	if err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	db := database.DbManager()
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	err = db.Collection("users").FindOne(ctx, bson.M{"_id": objectId}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		return c.Status(404).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	} else if err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&model.Json{
		Status: "success",
		Data:   u,
	})
}
