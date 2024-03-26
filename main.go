// package main

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	// "go.mongodb.org/mongo-driver/mongo/options"
// 	// "go.mongodb.org/mongo-driver/mongo/primitive"
// 	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
// )

// type Mongoinstance struct {
// 	Client *mongo.Client
// 	Db     *mongo.Database
// }

// var mg Mongoinstance

// const dbName = "go_project"

// const mongoUrl = "mongodb+srv://vinayakrajqaz:password.1ahvx7b.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

// type Employee struct {
// 	Id      string  `json:"id,omitempty" bson:"_id,omitempty"`
// 	Name    string  `json:"name"`
// 	Salarly float64 `json:"salarly"`
// 	Age     float64 `json:"age"`
// }

// func Connect() {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	err = client.Connect(ctx)
// 	db := client.Database(dbName)

// 	if err != nil {
// 		return err
// 	}

// 	mg = Mongoinstance{
// 		Client: client,
// 		Db:     db,
// 	}

// }

// func main() {

// 	if err := Connect(); err != nil {
// 		log.Fatal(err)
// 	}

// 	app := fiber.New()

// 	app.Get("/employee", func(c *fiber.Ctx) error {
// 		query := bson.D{{}}
// 		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
// 		if err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}
// 		var employee []Employee = make([]Employee, 0)

// 		if err := cursor.All(c.Context(), &employee); err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}
// 		return c.JSON(employee)
// 	})
// 	app.Post("/employee", func(c *fiber.Ctx) error {
// 		collection := mg.Db.Collection("employees")
// 		employee := new(Employee)
// 		if err := c.BodyParser(employee); err != nil {
// 			return c.Status(400).SendString(err.Error())
// 		}

// 		employee.Id = ""
// 		insertionResult, err := collection.InsertOne(c.Context(), employee)
// 		if err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}
// 		filter := bson.D{{key: "_id", value: insertionResult.InsertedID}}
// 		createdRecord := collection.findOne(c.Context(), filter)
// 		createdEmployee := &Employee{}
// 		createdRecord.Decode(createdEmployee)

// 		return c.Status(201).JSON(createdEmployee)
// 	})
// 	app.Put("/employee:id")
// 	app.Delete("/employee:id")

// }
// package main

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Mongoinstance struct {
// 	Client *mongo.Client
// 	Db     *mongo.Database
// }

// var mg Mongoinstance

// const dbName = "go_project"

// const mongoUrl = "mongodb+srv://vinayakrajqaz:password.1ahvx7b.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

// type Employee struct {
// 	Id      string  `json:"id,omitempty" bson:"_id,omitempty"`
// 	Name    string  `json:"name"`
// 	Salary  float64 `json:"salary"`
// 	Age     float64 `json:"age"`
// }

// func Connect() error {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
// 	if err != nil {
// 		return err
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	err = client.Connect(ctx)
// 	db := client.Database(dbName)

// 	if err != nil {
// 		return err
// 	}

// 	mg = Mongoinstance{
// 		Client: client,
// 		Db:     db,
// 	}

// 	return nil
// }

// func main() {
// 	err := Connect()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	app := fiber.New()

// 	app.Get("/employee", func(c *fiber.Ctx) error {
// 		query := bson.D{{}}
// 		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
// 		if err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}
// 		defer cursor.Close(c.Context())

// 		var employees []Employee
// 		if err := cursor.All(c.Context(), &employees); err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}
// 		return c.JSON(employees)
// 	})

// 	app.Post("/employee", func(c *fiber.Ctx) error {
// 		employee := new(Employee)
// 		if err := c.BodyParser(employee); err != nil {
// 			return c.Status(400).SendString(err.Error())
// 		}

// 		result, err := mg.Db.Collection("employees").InsertOne(c.Context(), employee)
// 		if err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}

// 		employee.Id = result.InsertedID.(primitive.ObjectID).Hex()
// 		return c.Status(201).JSON(employee)
// 	})

// 	log.Fatal(app.Listen(":3000"))
// }

// package main

// import (
// 	"context"
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Mongoinstance struct {
// 	Client *mongo.Client
// 	Db     *mongo.Database
// }

// var mg Mongoinstance

// const dbName = "go_project"

// const mongoUrl = "mongodb+srv://vinayakrajqaz:password.1ahvx7b.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

// type Employee struct {
// 	Id     string  `json:"id,omitempty" bson:"_id,omitempty"`
// 	Name   string  `json:"name"`
// 	Salary float64 `json:"salary"`
// 	Age    float64 `json:"age"`
// }

// func Connect() error {
// 	clientOptions := options.Client().ApplyURI(mongoUrl)
// 	client, err := mongo.Connect(context.Background(), clientOptions)
// 	if err != nil {
// 		return err
// 	}

// 	db := client.Database(dbName)

// 	mg = Mongoinstance{
// 		Client: client,
// 		Db:     db,
// 	}

// 	return nil
// }

// func main() {
// 	err := Connect()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	app := fiber.New()

// 	app.Get("/employee", func(c *fiber.Ctx) error {
// 		query := bson.D{{}}
// 		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
// 		if err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}
// 		defer cursor.Close(c.Context())

// 		var employees []Employee
// 		if err := cursor.All(c.Context(), &employees); err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}
// 		return c.JSON(employees)
// 	})

// 	app.Post("/employee", func(c *fiber.Ctx) error {
// 		employee := new(Employee)
// 		if err := c.BodyParser(employee); err != nil {
// 			return c.Status(400).SendString(err.Error())
// 		}

// 		result, err := mg.Db.Collection("employees").InsertOne(c.Context(), employee)
// 		if err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}

// 		employee.Id = result.InsertedID.(primitive.ObjectID).Hex()
// 		return c.Status(201).JSON(employee)
// 	})
// 	app.Put("/employee/:id", func(c *fiber.Ctx) error {
// 		idParam := c.Params("id")
// 		employeeID, err := primitive.ObjectIDFromHex(idParam)

// 		if err != nil {
// 			return c.SendStatus(400)
// 		}

// 		employee := new(Employee)

// 		if err := c.BodyParser(employee); err != nil {
// 			return c.Status(400).SendString(err.Error())
// 		}

// 		query := bson.D{{key: "_id", Value: employeeID}}
// 		update := bson.D{
// 			{Key: "$set",
// 				Value: bson.D{
// 					{Key: "name", Value: employee.Name},
// 					{Key: "age", Value: employee.Age},
// 					{Key: "salarly", Value: employee.Salary},
// 				},
// 			},
// 		}
// 		err :=mg.Db.Collection("employee").FindOneAndUpdate(c.Context(),query,update);

// 		if err != nil {
// 			if err==mongo.ErrNoDocuments {
// 				return c.SendStatus(400)
				
// 			}
// 			return c.SendStatus(500)
// 		}
// 		employee.Id=idParam
// 		return c.Status(200).JSON(employee)

// 	})
// 	app.Delete("/employee/:id",func(c *fiber.Ctx) error {
// 		employeeID, err :=primitive.ObjectIDFromHex(c.Params("id"))

// 		if err != nil {
// 			return c.SendStatus(400)
// 		}

// 		query := bson.D{{key: "_id", Value: employeeID}}
// 		result,err:=mg.Db.Collection("employee").DeleteOne(c.Context(),&query,)

// 		if err != nil {
// 			return c.SendStatus(500)
// 		}

// 		if result.DeletedCount<1{
// 			return c.SendStatus(404)
// 		}
// 		return c.Status(200).JSON("Record Deleated")
// 	})


// 	log.Fatal(app.Listen(":3000"))
// }



// package main

// import (
// 	"context"
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Mongoinstance struct {
// 	Client *mongo.Client
// 	Db     *mongo.Database
// }

// var mg Mongoinstance

// const dbName = "go_project"

// const mongoUrl = "mongodb+srv://vinayakrajqaz:C2PGMd4NPmXnDYrk@cluster0@cluster0.1ahvx7b.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

// type Employee struct {
// 	Id     string  `json:"id,omitempty" bson:"_id,omitempty"`
// 	Name   string  `json:"name"`
// 	Salary float64 `json:"salary"`
// 	Age    float64 `json:"age"`
// }

// func Connect() error {
// 	clientOptions := options.Client().ApplyURI(mongoUrl)
// 	client, err := mongo.Connect(context.Background(), clientOptions)
// 	if err != nil {
// 		return err
// 	}

// 	db := client.Database(dbName)

// 	mg = Mongoinstance{
// 		Client: client,
// 		Db:     db,
// 	}

// 	return nil
// }

// func main() {
// 	err := Connect()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	app := fiber.New()

// 	app.Get("/employee", func(c *fiber.Ctx) error {
// 		query := bson.D{{}}
// 		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
// 		if err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}
// 		defer cursor.Close(c.Context())

// 		var employees []Employee
// 		if err := cursor.All(c.Context(), &employees); err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}
// 		return c.JSON(employees)
// 	})

// 	app.Post("/employee", func(c *fiber.Ctx) error {
// 		employee := new(Employee)
// 		if err := c.BodyParser(employee); err != nil {
// 			return c.Status(400).SendString(err.Error())
// 		}

// 		result, err := mg.Db.Collection("employees").InsertOne(c.Context(), employee)
// 		if err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}

// 		employee.Id = result.InsertedID.(primitive.ObjectID).Hex()
// 		return c.Status(201).JSON(employee)
// 	})

// 	app.Put("/employee/:id", func(c *fiber.Ctx) error {
// 		idParam := c.Params("id")
// 		employeeID, err := primitive.ObjectIDFromHex(idParam)
// 		if err != nil {
// 			return c.SendStatus(400)
// 		}

// 		employee := new(Employee)
// 		if err := c.BodyParser(employee); err != nil {
// 			return c.Status(400).SendString(err.Error())
// 		}

// 		filter := bson.D{{Key: "_id", Value: employeeID}}
// 		update := bson.D{
// 			{Key: "$set", Value: bson.D{
// 				{Key: "name", Value: employee.Name},
// 				{Key: "age", Value: employee.Age},
// 				{Key: "salary", Value: employee.Salary},
// 			}},
// 		}
// 		_, err = mg.Db.Collection("employees").UpdateOne(c.Context(), filter, update)
// 		if err != nil {
// 			return c.SendStatus(500)
// 		}
// 		employee.Id = idParam
// 		return c.Status(200).JSON(employee)
// 	})

// 	app.Delete("/employee/:id", func(c *fiber.Ctx) error {
// 		employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))
// 		if err != nil {
// 			return c.SendStatus(400)
// 		}

// 		filter := bson.D{{Key: "_id", Value: employeeID}}
// 		result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), filter)
// 		if err != nil {
// 			return c.SendStatus(500)
// 		}

// 		if result.DeletedCount < 1 {
// 			return c.SendStatus(404)
// 		}
// 		return c.Status(200).JSON("Record Deleted")
// 	})

// 	log.Fatal(app.Listen(":3000"))
// }


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

const mongoUrl = "mongodb+srv://vinayakrajqaz:2D21NHXScPrOMp2p@cluster0.gucjl.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

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
