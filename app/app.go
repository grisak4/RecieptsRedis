package app

import (
	"fmt"
	"reciepts/config"
	"reciepts/config/getconf/getservconf"
	"reciepts/database"
	"reciepts/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	config.InitConfig()

	servconf := getservconf.GetServConfig()
	r := gin.Default()

	database.StartDB()
	defer database.CloseDB()

	routes.InitRoutes(r, database.GetDB())

	r.Run(fmt.Sprintf("%s:%d",
		servconf.Host, servconf.Port))
}
