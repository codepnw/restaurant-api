package controllers

import (
	"github.com/codepnw/restaurant-api/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")