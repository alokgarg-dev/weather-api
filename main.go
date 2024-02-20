package main

import (
	"context"
	"fmt"
	// "weather-api/application"
	"github.com/alokgarg-dev/weather-api/application"

	"github.com/spf13/viper"

	// c "github.com/alokgarg-dev/weather-api/config"
)

func main() {
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	// viper.AutomaticEnv()

	viper.SetConfigType("yml")
	// var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	fmt.Println("API_KEYis\t", viper.GetString("API_KEY"))

	/*
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("API_KEY is\t", configuration.ApiKey)
	*/

	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
