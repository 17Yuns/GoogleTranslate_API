package main

import (
	router "ApiController/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("Env Error:", err)
	}
	port := os.Getenv("Port")
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	router.DefaultRouterInit(r)
	PortErr := r.Run(":" + port)
	if PortErr != nil {
		logrus.Error("Port Error:", PortErr)
	}
}
