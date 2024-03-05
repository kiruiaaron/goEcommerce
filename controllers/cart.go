package controllers

import (
	"Context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Appplication struct{
	prodCollection * mongo.Collection
	userCollection *mongo.Collection
}


func NewApplication(prodCollection,userCollection *mongo.Collection)*Appplication{return &Appplication{
	prodCollection: prodCollection,
	userCollection: userCollecton,

}}

func  (app *Appplication) AddToCart() gin.Handler{

	return func (c *gin.Context)  {
		productQueryID := c.query("id")
		if productQueryID ==""{
			log.Println("Product is is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product id is empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == ""{
			log.Println("user id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
		}
	}




}

func RemoveItem()gin.HandlerFunc{

}

func GetItemFromCart() gin.HandlerFunc{

}

func BuyFromCart() gin.HandlerFunc{

}

func InstantBuy() gin.HandlerFunc{

}
