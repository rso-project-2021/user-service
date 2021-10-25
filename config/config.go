package config

import "github.com/spf13/viper"

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	GinMode       string `mapstructure:"GIN_MODE"`
}

var config Config

func Load(path string) (err error) {

	// set viper parameters
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// read env variables
	if err = viper.ReadInConfig(); err != nil {
		return
	}

	// save env variables into config object
	err = viper.Unmarshal(&config)
	return
}

func Read() Config {
	return config
}
