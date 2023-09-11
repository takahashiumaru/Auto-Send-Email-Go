package configuration

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	AccessSecret  string `mapstructure:"ACCESS_SECRET"`
	RefreshSecret string `mapstructure:"REFRESH_SECRET"`
	Port          string `mapstructure:"PORT"`
	PortDBMS      string `mapstructure:"PORT_DB_MS"`
	HostMS        string `mapstructure:"HOST_DB_MS"`
	PasswordMS    string `mapstructure:"PASSWORD_DB_MS"`
	UserMS        string `mapstructure:"USER_DB_MS"`
	DbMS          string `mapstructure:"DATABASE_DB_MS"`
	PortDBMY      string `mapstructure:"PORT_DB_MY"`
	HostMY        string `mapstructure:"HOST_DB_MY"`
	PasswordMY    string `mapstructure:"PASSWORD_DB_MY"`
	UserMY        string `mapstructure:"USER_DB_MY"`
	DbMY          string `mapstructure:"DATABASE_DB_MY"`
	FromEmail     string `mapstructure:"FROM_EMAIL"`
	PasswordEmail string `mapstructure:"PASSWORD_EMAIL"`
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
