package main

import (
	"log"
	"os"

	"github.com/Szym0nion/eCommerce-Go-Docker/controllers"
	"github.com/Szym0nion/eCommerce-Go-Docker/database"
	"github.com/Szym0nion/eCommerce-Go-Docker/middlewere"
	"github.com/Szym0nion/eCommerce-Go-Docker/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middlewere.Authentication())

	router.GET("addtocart", app.AddtoCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
