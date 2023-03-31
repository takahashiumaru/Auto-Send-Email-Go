package configuration

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	AccessSecret  string `mapstructure:"ACCESS_SECRET"`
	RefreshSecret string `mapstructure:"REFRESH_SECRET"`
	Port          string `mapstructure:"PORT"`
	PortDB        string `mapstructure:"PORT_DB"`
	Host          string `mapstructure:"HOST_DB"`
	Password      string `mapstructure:"PASSWORD_DB"`
	User          string `mapstructure:"USER_DB"`
	Db            string `mapstructure:"DATABASE_DB"`
}

func LoadConfig() (config Configuration, err error) {
	viper.SetConfigFile("./configuration/.env")
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
