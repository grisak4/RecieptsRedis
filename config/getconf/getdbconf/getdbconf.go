package getdbconf

import (
	"reciepts/config/confmodels"

	"github.com/spf13/viper"
)

func GetDBConfig() confmodels.DBconf {
	return confmodels.DBconf{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DBname:   viper.GetString("database.dbname"),
	}
}
