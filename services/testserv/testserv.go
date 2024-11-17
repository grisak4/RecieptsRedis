package testserv

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestSmth(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
