package main

import (
	_ "auth/docs"
	"auth/init/db"
	"auth/init/config"
	"auth/routes"
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	db.InitDB(os.Getenv("DB_URL"))
	fmt.Println(os.Getenv("DB_URL"))
	routes.SetupRoutes(router)
	router.Run(os.Getenv("PORT"))
}
