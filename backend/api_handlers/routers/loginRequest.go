package routers

import (
	"context"
	"fyndguru-hackathon/backend/models"
	"fyndguru-hackathon/backend/mongopack"
	"fyndguru-hackathon/backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func Login(c *gin.Context) {
	var u models.LoginReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:

	collection := mongopack.MongoDb.Collection("user_login")

	background := context.Background()
	cursor, _ := collection.Find(
		background,
		bson.D{{"userName", u.Username}})

	if cursor.Next(background) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Please provide valid login details")
			return
		}
		p := result["password"]
		if nil == p {
			c.JSON(http.StatusUnauthorized, "User password is not associated with us")
			return
		}

		doesMatch := utils.CheckPasswordHash(u.Password, p.(string))

		if !doesMatch {
			c.JSON(http.StatusUnauthorized, "Please provide valid login details")
			return
		}

		userId := result["userId"].(string)

		token, err := utils.CreateJwtToken(userId)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		c.JSON(http.StatusOK, token)

	}
}

func EmployerLogin(c *gin.Context) {
	var u models.LoginReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:

	collection := mongopack.MongoDb.Collection("employer_login")
	background := context.Background()
	cursor, _ := collection.Find(
		background,
		bson.D{{"userName", u.Username}})

	if cursor.Next(background) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Please provide valid login details")
			return
		}
		p := result["password"]
		if nil == p {
			c.JSON(http.StatusUnauthorized, "Employer password is not associated with us")
			return
		}

		doesMatch := utils.CheckPasswordHash(u.Password, p.(string))

		if !doesMatch {
			c.JSON(http.StatusUnauthorized, "Please provide valid login details")
			return
		}

		userId := result["userId"].(string)

		token, err := utils.CreateJwtToken(userId)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		c.JSON(http.StatusOK, token)

	}
}
