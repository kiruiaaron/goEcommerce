package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kiruiaaron/goEcommerce/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		address, err := primitive.ObjectIDFromHex(user_id)
		if err != nil{
			c.IndentedJSON(500, "INternal server error")
		}
		var addresses models.Address

		addresses.Address_id=primitive.NewObjectID()
		if err = c.BindJSON(&addresses);err != nil{
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx ,cancel = context.WithTimeout(context.Background(), 100*time.Second)
		match_filter := bson.D{{Key:"$match",Value:bson.D{primitive.E{Key:"_id",Value:address}}}}
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
			update := bson.D{{Key:"$push",Value:bson.D{primitive.E{Key:"address",Value: address}}}}
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
	return func (c *gin.Context)  {
		user_id := c.Query("id")
		if user_id == ""{
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid"})
			c.Abort()
			return 
		}
		usert_id , err := primitive.ObjectIDFromHex(user_id)
		if err != nil{
			c.IndentedJSON(500, "Internal server error")
		}

		var editAddress models.Address
        if err := c.BindJSON(&editAddress); err != nil{
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key:"_id", Value:usert_id}}
		update := bson.D{{Key:"$set",Value:bson.D{primitive.E{Key:"address.0.house_name", Value: editAddress.House},{Key: "address.0.street_name",Value: editAddress.Street},{Key: "address.0.city_name", Value: editAddress.City},{Key: "address.0.pin_code", Value: editAddress.Pincode}}}}
		_, err = userCollection.UpdateOne(ctx, filter, update)
		if err != nil{
			c.IndentedJSON(500, "something went wrong")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Successfully updated the home address")
	}

}


func EditWorkAddress() gin.HandlerFunc{
	return func (c *gin.Context)  {
		
	}

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
		filter := bson.D{primitive.E{Key:"_id", Value:usert_id}}
		update := bson.D{{Key:"$set",Value:bson.D{primitive.E{Key:"addresses",Value:addresses}}}}
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