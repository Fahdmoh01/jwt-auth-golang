package controllers

import (
	"jwt-auth-golang/initializers"
	"jwt-auth-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func Signup(c *gin.Context){
	// Get the email/pass off request body
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Failed to read body",
		})

		return 
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hush password",
		})

		return
	}
	//Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result :=  initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	//Response
	c.JSON(http.StatusOK, gin.H{})
}