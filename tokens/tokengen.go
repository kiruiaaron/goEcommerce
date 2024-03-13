package tokens

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kiruiaaron/goEcommerce/database"
	"go.mongodb.org/mongo-driver/mongo"
)


type SignedDetails struct{
	Email string 
	First_Name string
	Last_Name string
	Uid string
	jwt.StandardClaims
}


var UserData *mongo.Collection = database.UserData(database.Client,"Users")

var SECRET_KEY = os.Getenv("SECRET_KEY")

func TokenGenerator(email string, firstName string, lastName string, uid string)(signedToken string, signedRefreshToken string, err error){

	claims := &SignedDetails{
		Email: email,
		First_Name: firstName,
		Last_Name: lastName,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},

	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil{
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS384, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil{
		log.Panic(err)
		return

    }

	return token, refreshToken, err


}


func ValidateToken(){

}


func UpdateAllTokens(){

}