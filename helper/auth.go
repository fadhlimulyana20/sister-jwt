package helper

import (
	"context"
	"errors"
	"strings"

	"github.com/fadhlimulyana20/sister-jwt/database"
	"github.com/fadhlimulyana20/sister-jwt/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct{}

var ctx = context.Background()

func (a *Auth) GetUser(c *fiber.Ctx) (model.User, error) {
	authHeader := c.Get("Authorization")
	u := new(model.User)

	if !strings.Contains(authHeader, "Bearer") {
		return *u, errors.New("Belum Login")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	j := new(JWT)
	userId, err := j.Parse(tokenString, "token")

	if err != nil {
		return *u, err
	}

	db := database.DbManager()
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return *u, err
	}

	err = db.Collection("users").FindOne(ctx, bson.M{"_id": objectId}).Decode(&u)
	if err != nil {
		return *u, err
	}

	return *u, nil
}
