package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	//get the email/pass off req body
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Hash the password 
	
}
