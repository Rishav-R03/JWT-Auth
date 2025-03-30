package main

import (
	"22BSA10076_Backend/initializers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()  // ✅ Connect to DB using GORM
	initializers.SyncDatabase() // ✅ AutoMigrate models
}

func main() {
	fmt.Println("Welcome to JWT/Server!")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Home",
		})
	})

	r.Run()
}
