package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongoinstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg Mongoinstance

const dbName = "go_project"

const mongoUrl = "MONGODB_URL"

type Employee struct {
	Id     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

func Connect() error {
	clientOptions := options.Client().ApplyURI(mongoUrl)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	db := client.Database(dbName)

	mg = Mongoinstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func main() {
	err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {
		query := bson.D{{}}
		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer cursor.Close(c.Context())

		var employees []Employee
		if err := cursor.All(c.Context(), &employees); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(employees)
	})

	app.Post("/employee", func(c *fiber.Ctx) error {
		employee := new(Employee)
		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		result, err := mg.Db.Collection("employees").InsertOne(c.Context(), employee)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		employee.Id = result.InsertedID.(primitive.ObjectID).Hex()
		return c.Status(201).JSON(employee)
	})

	app.Put("/employee/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		employeeID, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			return c.SendStatus(400)
		}

		employee := new(Employee)
		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: employeeID}}
		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			}},
		}
		_, err = mg.Db.Collection("employees").UpdateOne(c.Context(), filter, update)
		if err != nil {
			return c.SendStatus(500)
		}
		employee.Id = idParam
		return c.Status(200).JSON(employee)
	})

	app.Delete("/employee/:id", func(c *fiber.Ctx) error {
		employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.SendStatus(400)
		}

		filter := bson.D{{Key: "_id", Value: employeeID}}
		result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), filter)
		if err != nil {
			return c.SendStatus(500)
		}

		if result.DeletedCount < 1 {
			return c.SendStatus(404)
		}
		return c.Status(200).JSON("Record Deleted")
	})

	log.Fatal(app.Listen(":3000"))
}
