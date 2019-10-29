package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

// ReadDBConfig would return connection string for database
func ReadDBConfig() *mgo.DialInfo {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	viper.SetConfigType("toml")

	uri, ok := viper.Get("mongo_uri").(string)
	if !ok {
		panic("database url is invalid")
	}

	dbConfig := &mgo.DialInfo{}

	viper.Unmarshal(dbConfig)
	dbConfig.Addrs = []string{uri}

	return dbConfig
}
