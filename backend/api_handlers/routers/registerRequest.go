package routers

import (
	"context"
	"fyndguru-hackathon/backend/models"
	"fyndguru-hackathon/backend/mongopack"
	"fyndguru-hackathon/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var u models.UserRegistration
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	loginCollection := mongopack.MongoDb.Collection("user_login")
	verifyRegistration(c, loginCollection, u)
}

func EmployerRegister(c *gin.Context) {
	var u models.UserRegistration
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	loginCollection := mongopack.MongoDb.Collection("employer_login")
	verifyRegistration(c, loginCollection, u)
}

func verifyRegistration(c *gin.Context, loginCollection *mongo.Collection, u models.UserRegistration) {
	cursor, err := loginCollection.Find(
		context.Background(),
		bson.D{{"userName", u.Username}},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Please provide valid login details")
			return
		}
		userName := result["userName"]

		if userName != nil {
			c.JSON(http.StatusBadRequest, "username already taken")
			return
		}
	}

	password, er := utils.HashPassword(u.Password)

	if er != nil {
		c.JSON(http.StatusBadRequest, "Unable to parse password")
		return
	}

	uid := uuid.New().String()
	u.UserId = uid
	u.Password = password

	_, _ = loginCollection.InsertOne(context.Background(), u)

	token, err := utils.CreateJwtToken(uid)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
