package controllers

import (
	"context"
	"fmt"
	"image/color/palette"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kiruiaaron/goEcommerce/models"
	"go.mongo.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress()gin.HandlerFunc{
	return func (c *gin.Context)  {
		user_id := c.Query("id")
		if user_id == ""{
			c.Header("Content-Type","application")
			c.JSON(http.StatusNotFound, gin.H{"error":"invalid code"})
			c.Abort()
			return

		}
		address, err := ObjectIDFromHex(user_id)
		if err != nil{
			c.IndentedJSON(500, "INternal server error")
		}
		var addresses models.Address

		addresses.Address_id=primitive.NewObjectID()
		if err = c.BindJSON(&addresses);err != nil{
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx ,cancel = context.WithTimeout(context.Background(), 100*time.Second)
		match_filter := bson.D{Key:"$match",Value:bson.D{primitive.E{Key:"_id",Value:address}}}
		unwind := bson.D{{Key:"$unwind", Value:bson.D{primitive.E{Key:"path",Value:"$address"}}}}
		group := bson.D{{Key:"$group", Value:bson.D{primitive.E{Key:"_id",Value:"$address_id"},{Key:"count",Value:bson.D{primitive.E{Key:"$sum", Value:1}}}}}}
		pointcursor, err := userCollection.Aggregate(ctx, mongo.Pipeline{match_filter,unwind,group})

		if err != nil{
			c.IndentedJSON(500, "Internal Server error")
		}

		var addressInfo []bson.M
		if err = pointcursor.All(ctx, &addressInfo); err != nil{
			panic(err)
		}

		var size int32
		for _, address_no := range addressInfo {
		count := address_no["count"]
		size = count.(int32)
		}
		if size < 2{
			filter := bson.D{primitive.E{Key:"_id", Value:address}}
			update := bson.D{{Key:"$push",Value:bson.D{primitive.E{key:"address",Value: address}}}}
			_, err := userCollection.UpdateOne(ctx, filter,update)
			if err != nil{
				fmt.Println(err)
			}
		    }else {
			c.IndentedJSON(400, "Not allowed")
		}
		defer cancel()

		ctx.Done()

	}

}

func EditHomeAddress() gin.HandlerFunc{

}


func EditWorkAddress() gin.HandlerFunc{

}

func DeleteAddress() gin.HandlerFunc{
	return func (c * gin.Context)  {
		user_id := c.Query("id")
		if user_id == ""{
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid Search Index"})
			c.Abort()
			return 
		}

		addresses := make([]models.Address, 0)
		usert_id , err := primitive.ObjectIDFromHex(user_id)
		if err != nil{
			c.IndentedJSON(500, "Internal server error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{key:"_id", Vlalue:usert_id}}
		update := bson.D{{key:"$set",Value:bson.D{primitive.E{key:"addresses",Value:addresses}}}}
		_, err = userCollection.UpdateOne(ctx,filter,update)
		if err != nil{
			c.IndentedJSON(404, "wrong command")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Successfully Deleted")
	}
}