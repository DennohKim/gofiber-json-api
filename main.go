package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/DennohKim/gofiber-json-api/api"
	"github.com/DennohKim/gofiber-json-api/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"
const dbName = "gofiber-json-api"
const userColl = "users"

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	coll := client.Database(dbName).Collection(userColl)

	user := types.User{
		FirstName: "James",
		LastName:  "Bond",
	}

	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)

	}

	var james types.User

	if err :=  coll.FindOne(ctx, bson.M{}).Decode(&james); err != nil {
		log.Fatal(err)
	}

	fmt.Println(james)

	listenAddr := flag.String("listenAddr", ":5000", "API server listen address")
	flag.Parse()
	app := fiber.New()

	apiv1 := app.Group("api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}
