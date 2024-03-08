package controllers

import (
	"Context"
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kiruiaaron/goEcommerce/database"
	"github.com/kiruiaaron/goEcommerce/models"
	"go.mongo.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Appplication struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Appplication {
	return &Appplication{
		prodCollection: prodCollection,
		userCollection: userCollecton,
	}
}

func (app *Appplication) AddToCart() gin.Handler {

	return func(c *gin.Context) {
		productQueryID := c.query("id")
		if productQueryID == "" {
			log.Println("Product is is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product id is empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return

		}
		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return

		}
		c.IndentedJSON(200, "sucessfully added to the cart ")
	}

}

func (app *Appplication) RemoveItem() gin.HandlerFunc {

	return func(c *gin.Context) {
		productQueryID := c.query("id")
		if productQueryID == "" {
			log.Println("Product is is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product id is empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return

		}
		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return

		}
		c.IndentedJSON(200, "sucessfully remove the item from the cart ")

	}
}

func GetItemFromCart() gin.HandlerFunc {

	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			c.Abort()
			return
		}

		usert_id, _ := primitive.ObjectIDFromHex(user_id)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		var filledCart models.User
		err := userCollection.FindOne(ctx, bson.D{primitive.E{key: "_id", Value: usert_id}}).Decode(&filledCart)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(500, "not found")
			return
		}

		filter_match := bson.D{key: "$match", Value: bson.D{primitive.E{key: "_id", Value: usert_id}}}
		unwind := bson.D{{key: "$unwind", Value: bson.D{primitive.E{key: "path", Value: "$usercart"}}}}
		grouping := bson.D{{key: "$group", Value: bson.D{primitive.E{key: "_id", Value: "$_id"}, key: "total", Value: bson.D{primitive.E{key: "$sum", Value: "$usercart.price"}}}}}

		pointCursor, err := userCollection.Aggregate(ctx, mongo.Pipeline{filter_match, unwind,grouping})
		if err != nil{
			log.Println(err)
		}
		var listing []bson.M
		if err = pointCursor.All(ctx, &listing);err != nil{
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		for _, json := range listing{
			c.IndentedJSON(200, json["total"])
			c.IndentedJSON(200, filledCart.UserCart)
		}
		ctx.Done()


	}

}

func (app *Appplication) BuyFromCart() gin.HandlerFunc {

	return func(c *gin.Context) {
		userQuery := c.Query("id")

		if userQuery == "" {
			log.Println("use id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("userID is empty"))
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		err := database.BuyFromCart(ctx, app.userCollection, userQuery)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)

		}

		c.IndentedJSON("successfully placed the order")

	}

}

func (app *Appplication) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.query("id")
		if productQueryID == "" {
			log.Println("Product is is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product id is empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return

		}
		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)

		}

		c.IndentedJSON(200, "successfully placed the order")

	}

}
