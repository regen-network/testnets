package db

import (
	"fmt"

	"github.com/spf13/viper"
)

// ReadDBConfig would return connection string for database
func ReadDBConfig() string {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	viper.SetConfigType("toml")

	mongoURI, ok := viper.Get("mongo_uri").(string)
	if !ok {
		panic("database url is invalid")
	}

	return mongoURI
}
