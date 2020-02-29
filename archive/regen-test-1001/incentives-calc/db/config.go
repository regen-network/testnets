package db

import (
	"fmt"

	"errors"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2"
)

type Config struct {
	Mongo_uri string `json:"mongo_uri" validate:"required"`
	Database  string `json:"database" validate:"required"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Source    string `json:"source"`
	FailFast  string `json:"failFast" validate:"required"`
}

// ReadDBConfig would return connection string for database
func ReadDBConfig() *mgo.DialInfo {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Errorf("fatal error config file: %s", err)
		HandleError(err)
	}

	viper.SetConfigType("toml")

	uri, ok := viper.Get("mongo_uri").(string)
	if !ok {
		//panic("database url is invalid")
		HandleError(errors.New("Database url is invalid"))
	}

	dbConfig := &mgo.DialInfo{}

	viper.Unmarshal(dbConfig)
	dbConfig.Addrs = []string{uri}

	cfg := &Config{}
	viper.Unmarshal(cfg)

	//Validating all required fields from config
	validate := validator.New()
	err = validate.Struct(cfg)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		fmt.Printf("Error Env variable are missing %v", validationErrors.Error())

		os.Exit(1)
	}

	return dbConfig
}

func HandleError(err error) {
	fmt.Printf("Error %v", err)
	os.Exit(1)
}
