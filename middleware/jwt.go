package middleware

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

func JwtMiddleware(c *fiber.Ctx) error {
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

	return c.Next()
}
