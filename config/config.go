package config

import (
	"log"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"db_driver"`
	DBSource      string `mapstructure:"db_source"`
	ServerAddress string `mapstructure:"server_address"`
	GinMode       string `mapstructure:"gin_mode"`
}

// Reads configuration from file or environment variables.
func New(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	// Some small change
	err = viper.Unmarshal(&config)

	return
}

func KeyWatcher(key string, handler func(source string)) {

	// Get consul client.
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	kv := client.KV()

	// Read initial db source if exists.
	pair, meta, err := kv.Get(key, nil)
	if err != nil {
		log.Panic("Unable to list keys", err)
	}

	if pair != nil {
		handler(string(pair.Value))
	}

	options := api.QueryOptions{
		RequireConsistent: true,
	}

	// In case consul configuration changes update
	// database connection.
	for {
		options.WaitIndex = meta.LastIndex
		pair, meta, err = kv.Get(key, &options)

		if err != nil {
			log.Panic("Unable to get key.", err)
			return
		}

		if pair != nil {
			handler(string(pair.Value))
		}
	}
}
