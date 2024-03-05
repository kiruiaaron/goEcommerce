package main

import (
	"os"

	"github.com/docker/docker/api/server/router"
	"github.com/gin-gonic/gin"
	"github.com/kiruiaaron/goEcommerce/controllers"
	"github.com/kiruiaaron/goEcommerce/database"
	"github.com/kiruiaaron/goEcommerce/middleware"
	"github.com/kiruiaaron/goEcommerce/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.use(middleware.Authentication())

	router.Get("/addtocart", app.AddToCart())
	router.Get("/removeitem", app.RemoveItem())
	router.Get("/cartcheckout", app.BuyFromCart())
	router.Get("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
