package main

import (
	"fmt"
	"log"
	"os"

	"golang-restaurant-management/database"
	"golang-restaurant-management/middlewares"

	"golang-restaurant-management/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var port string = os.Getenv("PORT")
	fmt.Printf("Port: %s\n", port)

	if port == "" {
		port = "8000"
	}
	fmt.Println("------------")
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middlewares.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)

	router.Run(":" + port)

}
