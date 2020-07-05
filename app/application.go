package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func init() {
	godotenv.Load()
}

// StartApplication ..
func StartApplication() {
	mapUrls()
	router.Run(":" + os.Getenv("PORT"))
}
