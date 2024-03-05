package models

import (
	"time"

	"github.com/go-faster/city"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

 type User struct{
	
	ID
	First_Name
	Last_Name
	Password
	Email
	Phone
	Token
	Refresh_Token
	Created_At
	Updated_At
	User_ID
	UserCart
	Address_Details
	Order_Status
 }

type Product struct{

	Product_ID
	Product_Name
	Price
	Rating
	Image
    
}

type ProductUser struct{

	Product_ID
	Product_Name
	Price
	Rating
	Image

}

type Address struct{

	Address_id
	House
	Street
	City
	Pincode

}

type Order struct{

	Order_ID
	Order_Cart
	Ordered_At
	Price
    Discount
	Payment_Method

}

type Payment struct{

	Digital bool
	COD bool

}