package routes

import (
	"reciepts/middlewares/cors"
	"reciepts/services/testserv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	cors.SetCors(router)

	router.GET("/hello", func(c *gin.Context) {
		testserv.TestSmth(c, db)
	})
}
