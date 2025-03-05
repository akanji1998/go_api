package main

import (
	"log"
	"os"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main()  {
	db.InitDB()
	server := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.RegisterRoute(server)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := server.Run(":" + port); err != nil {
        log.Panicf("error: %s", err)
	}
// server.Run(":8080") //localhost:8080
}
