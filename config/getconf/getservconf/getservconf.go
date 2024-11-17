package getservconf

import (
	"reciepts/config/confmodels"

	"github.com/spf13/viper"
)

func GetServConfig() confmodels.ServConf {
	return confmodels.ServConf{
		Host: viper.GetString("server.host"),
		Port: viper.GetInt("server.port"),
	}
}
