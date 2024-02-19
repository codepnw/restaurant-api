package controllers

import (
	"github.com/codepnw/restaurant-api/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")
