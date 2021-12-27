package config

import (
	"log"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

type Config struct {
	LogitAddress  string `mapstructure:"logit_address"`
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

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	// Overwrite with consul value.
	consulConfig := api.DefaultConfig()
	client, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	kv := client.KV()

	pair, _, err := kv.Get("db_source", nil)
	if err == nil && pair != nil {
		config.DBSource = string(pair.Value)
	}

	err = nil
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
	_, meta, err := kv.Get(key, nil)
	if err != nil {
		return
	}

	options := api.QueryOptions{
		RequireConsistent: true,
	}

	// In case consul configuration changes update
	// database connection.
	var pair *api.KVPair
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
