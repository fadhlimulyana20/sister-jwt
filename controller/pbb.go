package controller

import (
	"github.com/fadhlimulyana20/sister-jwt/database"
	"github.com/fadhlimulyana20/sister-jwt/helper"
	"github.com/fadhlimulyana20/sister-jwt/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePbb(c *fiber.Ctx) error {
	auth := new(helper.Auth)

	u, err := auth.GetUser(c)
	if err != nil {
		return c.Status(401).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	PBB := new(model.Pbb)
	if err := c.BodyParser(PBB); err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	PBB.Id = primitive.NewObjectID()
	PBB.UserID = u.Id
	PBB.DataPbb = []model.DataPBB{}

	db := database.DbManager()
	_, err = db.Collection("pbb").InsertOne(ctx, PBB)
	if err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&model.Json{
		Status: "success",
		Data:   PBB,
	})
}

func AddDataPbb(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	auth := new(helper.Auth)

	u, err := auth.GetUser(c)
	if err != nil {
		return c.Status(401).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	PBB := new(model.Pbb)

	db := database.DbManager()
	err = db.Collection("pbb").FindOne(ctx, bson.M{"_id": objectID}).Decode(&PBB)
	if err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	if PBB.UserID.Hex() != u.Id.Hex() {
		return c.Status(401).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	dataPBB := new(model.DataPBB)
	if err := c.BodyParser(dataPBB); err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	dataPBB.Id = primitive.NewObjectID()
	if _, err := db.Collection("pbb").UpdateOne(ctx, bson.D{{"_id", objectID}}, bson.D{primitive.E{"$push", bson.D{{"data_pbb", dataPBB}}}}); err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&model.Json{
		Status: "success",
		Data:   PBB,
	})

}

func GetPbb(c *fiber.Ctx) error {
	auth := new(helper.Auth)

	u, err := auth.GetUser(c)
	if err != nil {
		return c.Status(401).JSON(&model.Json{
			Status:  "False",
			Message: err.Error(),
		})
	}

	db := database.DbManager()

	var PBB []model.Pbb
	cursor, err := db.Collection("pbb").Find(ctx, bson.M{"user_id": u.Id})
	if err != nil {
		return c.Status(500).JSON(&model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	for cursor.Next(ctx) {
		var row model.Pbb
		err := cursor.Decode(&row)

		if err != nil {
			return c.Status(500).JSON(&model.Json{
				Status:  "failed",
				Message: err.Error(),
			})
		}

		PBB = append(PBB, row)
	}

	return c.Status(200).JSON(&model.Json{
		Status: "success",
		Data:   PBB,
	})
}
